package main

import (
	"encoding/json"
	"fmt"
	"hermes/socket/config"
	"hermes/socket/epoll"
	"hermes/socket/httpclient"
	"hermes/socket/pool"
	"hermes/socket/redisclient"
	"hermes/socket/utils"
	"io"
	"log"
	"net"
	"syscall"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Message struct {
	SenderId  int    `json:"senderId"`
	ClanId    int    `json:"clanId"`
	ChannelId int    `json:"channelId"`
	Message   string `json:"message"`
	Time      int    `json:"time"`
}

var (
	myepoll      *epoll.SocketEpoll
	mypool       *pool.GoPool
	messageQueue chan *Message
)

func main() {
	messageQueue = make(chan *Message, 10)
	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	// Load environment variable
	env, err := config.GetEnv()
	if err != nil {
		log.Println(err)
	}
	// make epoll and pool
	myepoll, err := epoll.CreateEpoll()
	if err != nil {
		log.Panic(err)
	}
	mypool := pool.NewGoPool(env.TASK_QUEUE_SIZE, env.MAX_WORKER_NUM, env.INIT_WOKER_NUM)

	// Make redis client
	redis := redisclient.NewRedisClient(env)
	log.Println(redis)

	// Make authentication client
	authClient := httpclient.NewAuthenticationClient(
		fmt.Sprintf("%v%v", env.AUTH_SERVICE_HOST, env.WS_AUTH_PATH),
	)
	go WaitAndRead(myepoll, mypool, messageQueue)
	go WaitAndWrite(myepoll, mypool, messageQueue)
	// Main bussiness
	ln, err := net.Listen("tcp", env.APP_PORT)
	if err != nil {
		log.Fatal(err)
	}
	authUser := httpclient.User{}
	u := utils.GetUpgrader(authClient, &authUser)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}

		_, err = u.Upgrade(conn)
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		} else {
			myepoll.AddSocket(conn, &authUser)
		}
	}
}

// WaitAndRead wait for messages from epoll,
// and then write message to message queue.
// you can only write to messageQueue channel.
func WaitAndRead(epoll *epoll.SocketEpoll, p *pool.GoPool, messageQueue chan<- *Message) {
	for {
		connections, err := epoll.Wait()
		if err != nil {
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				log.Println("conn nil")
				continue
			}
			// ---------------------------------------
			// TODO: move these block code to inside queue task
			// to improve performance when deal with alot of user,
			// but it really dangerous to do so.
			// The epoll keep emit event unil you finish reading from the emitted connection.
			// Be careful here.
			connFd := utils.GetFdFromConnection(conn)
			log.Println("2. conn start read: ", connFd)
			header, err := ws.ReadHeader(conn)
			if err != nil {
				log.Println("read header error: ", err)
				epoll.RemoveSocket(conn)
			}
			log.Println("3. conn finish read: ", connFd)
			payload := make([]byte, header.Length)
			_, err = io.ReadFull(conn, payload)
			if err != nil {
				log.Println("read payload err: ", err)
				epoll.RemoveSocket(conn)
			}
			// --------------------------------------
			p.Queue(func() {
				// TODO: move to here
				if header.Masked {
					ws.Cipher(payload, header.Mask, 0)
				}

				// Reset the Masked flag, server frames must not be masked as
				// RFC6455 says.
				header.Masked = false
				var message Message
				if err := json.Unmarshal(payload, &message); err != nil {
					log.Println("json unmarshal: ", err)
					epoll.RemoveSocket(conn)
				}
				log.Println("7. read: ", message)
				messageQueue <- &message
			})
		}
	}
}

// WaitAndWrite waits for message from messageQueue channel,
// and then send these messages to corresponding clients.
// you can only read from messageQueue channel.
func WaitAndWrite(epoll *epoll.SocketEpoll, pool *pool.GoPool, messageQueue <-chan *Message) {
	for {
		mess := <-messageQueue
		log.Println("8. receive", mess)
		clan := mess.ClanId
		fdList := epoll.GetFDByClan(clan)
		log.Println("9. send to FDs: ", fdList, " clan ", clan)
		for _, fd := range fdList {
			conn := epoll.GetConnectionByFD(fd)
			if conn == nil {
				continue
			}

			w := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
			encoder := json.NewEncoder(w)

			if err := encoder.Encode(mess); err != nil {
				log.Println("encode message fail: ", err)
				if err := epoll.RemoveSocket(conn); err != nil {
					log.Println("remove connection fail: ", err)
				}
			}

			if err := w.Flush(); err != nil {
				log.Println("flush message fail: ", err)
				if err := epoll.RemoveSocket(conn); err != nil {
					log.Println("remove connection fail: ", err)
				}
			}
		}
	}
}

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
	var authRes httpclient.AuthRes
	u := utils.GetUpgrader(authClient, &authRes)
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
			myepoll.AddSocket(conn, 1)
		}
	}
}

// WaitAndRead wait for messages from epoll,
// and then write message to message queue.
// you can only write to messageQueue channel.
func WaitAndRead(epoll *epoll.SocketEpoll, pool *pool.GoPool, messageQueue chan<- *Message) {
	for {
		connections, err := epoll.Wait()
		if err != nil {
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				continue
			}
			pool.Queue(func() {
				buff, _, err := wsutil.ReadClientData(conn)
				if err != nil {
					if err := epoll.RemoveSocket(conn, 1); err != nil {
						log.Printf("Failed to remove %v", err)
					}
				}
				var message Message
				if err := json.Unmarshal(buff, &message); err != nil {
					log.Println("fail to unmarshal: ", err)
					epoll.RemoveSocket(conn, 1)
				}
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
		log.Println("receive message", mess)
		clan := mess.ClanId
		fdList := epoll.GetFDByClan(clan)
		log.Println("wait and write file descriptor list: ", fdList)
		for _, fd := range fdList {
			conn := epoll.GetConnectionByFD(fd)

			w := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
			encoder := json.NewEncoder(w)

			if err := encoder.Encode(mess); err != nil {
				log.Println("encode message fail: ", err)
				if err := epoll.RemoveSocket(conn, 1); err != nil {
					log.Println("remove connection fail: ", err)
				}
			}

			if err := w.Flush(); err != nil {
				log.Println("flush message fail: ", err)
			}
		}
	}
}

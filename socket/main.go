package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/socket/config"
	"hermes/socket/epoll"
	"hermes/socket/httpclient"
	"hermes/socket/redisclient"
	"hermes/socket/utils"
	"log"
	"net"
	"syscall"

	"github.com/gobwas/ws/wsutil"
)

type Message struct {
	SenderId  int    `json:"senderId"`
	ClanId    int    `json:"clanId"`
	ChannelId int    `json:"channelId"`
	Message   string `json:"message"`
	Time      int    `json:"time"`
}

func main() {
	epoll, err := epoll.CreateEpoll()
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

	// Make redis client
	redis := redisclient.NewRedisClient(env)
	log.Println(redis)

	// Make authentication client
	authClient := httpclient.NewAuthenticationClient(
		fmt.Sprintf("%v%v", env.AUTH_SERVICE_HOST, env.WS_AUTH_PATH),
	)
	// Main bussiness
	ln, err := net.Listen("tcp", env.APP_PORT)
	if err != nil {
		log.Println(err)
	}
	var authRes httpclient.AuthRes
	u := utils.GetUpgrader(authClient, &authRes)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		_, err = u.Upgrade(conn)
		if err != nil {
			log.Println(err)
		} else {
			epoll.AddSocket(conn, 1)
		}
	}
}

func WaitAndRead(epoll *epoll.SocketEpoll, redis *redisclient.RedisClient) {
	for {
		connections, err := epoll.Wait()
		if err != nil {
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			buff, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				if err := epoll.RemoveSocket(conn); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				conn.Close()
			}
			var message Message
			if err := json.Unmarshal(buff, &message); err != nil {
				log.Println(err)
				break
			}
			if err := redis.Publish("1", buff); err != nil {
				log.Println("publish", err)
				break
			}

		}
	}
}
func WaitAndWrite(epoll *epoll.SocketEpoll, redis *redisclient.RedisClient) {
	subcribe := redis.Subcribe("1")
	for {
		var ctx = context.Background()
		mess, err := subcribe.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("receive message", mess)
		buff := []byte(mess.Payload)
		var receivedMessage Message
		if err := json.Unmarshal(buff, &receivedMessage); err != nil {
			log.Println(err)
			break
		}
		clan := receivedMessage.ClanId
		fdList := epoll.GetFDByClan(clan)
		for _, fd := range fdList {
			conn := epoll.GetConnectionByFD(fd)
			if err := wsutil.WriteClientBinary(conn, buff); err != nil {
				if err := conn.Close(); err != nil {
					continue
				}
				continue
			}
		}
		// writer.Write([]byte(mess.Payload))
		// if err = writer.Flush(); err != nil {
		// 	break
		// }
	}
}

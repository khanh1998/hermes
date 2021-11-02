package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hermes/socket/config"
	"hermes/socket/httpclient"
	"hermes/socket/redisclient"
	"log"
	"net"
	"regexp"
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

func main() {
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
	u := ws.Upgrader{
		OnRequest: func(uri []byte) error {
			log.Println("on request: ", string(uri))
			tokenParam := "token=[A-Za-z0-9-_=]+.[A-Za-z0-9-_=]+.?[A-Za-z0-9-_.+/=]*"
			regex, err := regexp.Compile(tokenParam)
			if err != nil {
				return err
			}
			matches := regex.FindStringSubmatch(string(uri))
			if len(matches) == 0 {
				return errors.New("no token")
			}
			log.Println(matches[0][6:])
			res, err := authClient.AuthenticateWebsocket(matches[0][6:])
			authRes = res
			if err != nil {
				return err
			}
			return nil
		},
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		_, err = u.Upgrade(conn)
		if err != nil {
			log.Println(err)
		} else {
			go func() {
				defer conn.Close()

				var (
					state  = ws.StateServerSide
					reader = wsutil.NewReader(conn, state)
				)
				for {
					header, err := reader.NextFrame()
					if err != nil {
						break
					}
					if header.OpCode == ws.OpClose {
						break
					}

					buff := make([]byte, header.Length)
					reader.Read(buff)
					if err := redis.Publish("1", buff); err != nil {
						log.Println("publish", err)
						break
					}
				}
			}()
			go func() {
				defer conn.Close()
				var (
					state     = ws.StateServerSide
					writer    = wsutil.NewWriter(conn, state, ws.OpText)
					myAuthRes = authRes
				)
				log.Println(myAuthRes.Username)
				for {
					// Reset writer to write frame with right operation code.
					// writer.Reset(conn, state, header.OpCode)
					subcribe := redis.Subcribe("1")
					var ctx = context.Background()
					mess, err := subcribe.ReceiveMessage(ctx)
					if err != nil {
						log.Println(err)
						break
					}
					log.Println("receive message", mess)
					var receivedMessage Message
					if err := json.Unmarshal([]byte(mess.Payload), &receivedMessage); err != nil {
						log.Println(err)
						break
					}
					// if myAuthRes.ID != receivedMessage.SenderId {
					writer.Write([]byte(mess.Payload))
					// }
					// if _, err = io.Copy(writer, reader); err != nil {
					// handle error
					// }
					if err = writer.Flush(); err != nil {
						// handle error
						break
					}
				}
			}()
		}
	}
}

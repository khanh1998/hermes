package main

import (
	"errors"
	"hermes/socket/httpclient"
	"io"
	"log"
	"net"
	"regexp"

	"github.com/gobwas/ws"
)

func main() {
	authClient := httpclient.NewAuthenticationClient("http://localhost:4000/authentication/ws")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
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
			if err = authClient.AuthenticateWebsocket(matches[0][6:]); err != nil {
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
		}

		go func() {
			defer conn.Close()

			for {
				header, err := ws.ReadHeader(conn)
				if err != nil {

				}
				payload := make([]byte, header.Length)
				_, err = io.ReadFull(conn, payload)
				log.Println(payload)
				if err != nil {
					log.Println(err)
				}
				if header.Masked {
					ws.Cipher(payload, header.Mask, 0)
				}
				header.Masked = false
				if err := ws.WriteHeader(conn, header); err != nil {
					log.Println(err)
					break
				}
				if _, err := conn.Write(payload); err != nil {
					log.Println(err)
					break
				}
				if header.OpCode == ws.OpClose {
					return
				}
			}
		}()
	}
}

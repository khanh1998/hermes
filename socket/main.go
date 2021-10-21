package main

import (
	"log"
	"net"

	"github.com/gobwas/ws"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	u := ws.Upgrader{
		OnHeader: func(key, value []byte) (err error) {
			log.Printf("non-websocket header: %q=%q", key, value)
			return
		},
		OnRequest: func(uri []byte) error {
			log.Println("on request: ", string(uri))
			return nil
		},
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}

		_, err = u.Upgrade(conn)
		if err != nil {
			// handle error
		}
	}
}

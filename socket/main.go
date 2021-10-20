package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("new connection")
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Panic(err)
		}
		go func() {
			defer conn.Close()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					//log.Panic(err)
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					//log.Panic(err)
				}
			}
		}()
	}))

	if err != nil {
		log.Panic(err)
	}
	log.Println("websocket server listening on port 8080")
}

package utils

import (
	"errors"
	"hermes/socket/httpclient"
	"log"
	"net"
	"reflect"
	"regexp"

	"github.com/gobwas/ws"
)

func GetUpgrader(authClient *httpclient.AuthenticationClient, authRes *httpclient.User) *ws.Upgrader {
	return &ws.Upgrader{
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
			authRes.ID = res.ID
			authRes.Username = res.Username
			authRes.Clans = res.Clans
			if err != nil {
				return err
			}
			return nil
		},
	}
}

// kernel maintains a file descriptor for each connection,
// this function get fd number of the input connection.
func GetFdFromConnection(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}

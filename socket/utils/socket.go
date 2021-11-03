package utils

import (
	"errors"
	"hermes/socket/httpclient"
	"log"
	"regexp"

	"github.com/gobwas/ws"
)

func GetUpgrader(authClient *httpclient.AuthenticationClient, authRes *httpclient.AuthRes) *ws.Upgrader {
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
			authRes = res
			if err != nil {
				return err
			}
			return nil
		},
	}
}

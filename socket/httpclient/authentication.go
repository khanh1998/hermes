package httpclient

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type AuthenticationClient struct {
	endpoint string
}

func NewAuthenticationClient(endpoint string) *AuthenticationClient {
	return &AuthenticationClient{
		endpoint: endpoint,
	}
}

func (a *AuthenticationClient) AuthenticateWebsocket(token string) error {
	req, err := http.NewRequest(http.MethodPost, a.endpoint, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("got error here: ", err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusCreated {
		return nil
	}
	return errors.New("Invalid token")
}

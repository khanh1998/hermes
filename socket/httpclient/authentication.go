package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthenticationClient struct {
	endpoint string
}
type Clan struct {
	ID      int    `json:"id"`
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	ChiefId int    `json:"chiefId"`
}
type User struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	Clans    []struct {
		ID      int    `json:"id"`
		Domain  string `json:"domain"`
		Name    string `json:"name"`
		ChiefId int    `json:"chiefId"`
	} `json:"clans"`
}

func NewAuthenticationClient(endpoint string) *AuthenticationClient {
	return &AuthenticationClient{
		endpoint: endpoint,
	}
}

func (a *AuthenticationClient) AuthenticateWebsocket(token string) (*User, error) {
	req, err := http.NewRequest(http.MethodPost, a.endpoint, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("got error here: ", err)
		return &User{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var authRes User
	json.Unmarshal(body, &authRes)
	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return &authRes, nil
	}
	return &User{}, errors.New("Invalid token")
}

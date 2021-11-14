package esclient

import (
	"context"
	"encoding/json"
	"errors"
	"hermes/shipping/config"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	es7 "github.com/elastic/go-elasticsearch/v7"
)

type ElasticSearchClient struct {
	client *es7.Client
}

func NewElasticSearchClient(env *config.Env) (*ElasticSearchClient, error) {
	cfg := es7.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := es7.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	info, err := es.Info()
	log.Println(info)
	if err != nil {
		return nil, err
	}
	defer info.Body.Close()
	return &ElasticSearchClient{
		client: es,
	}, nil
}

func (e *ElasticSearchClient) SendMessage(message []byte) error {
	req := esapi.IndexRequest{
		Index:   "message",
		Body:    strings.NewReader(string(message)),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return errors.New(res.Status())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
		return nil
	}
}

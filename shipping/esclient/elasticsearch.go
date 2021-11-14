package esclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hermes/shipping/config"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/segmentio/kafka-go"
)

type ElasticSearchClient struct {
	client *es7.Client
}

func NewElasticSearchClient(env *config.Env) (*ElasticSearchClient, error) {
	cfg := es7.Config{
		Addresses: []string{
			env.ELASTIC_SEARCH_URI,
		},
	}
	es, err := es7.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	info, err := es.Info()
	if err != nil {
		return nil, err
	}
	defer info.Body.Close()
	return &ElasticSearchClient{
		client: es,
	}, nil
}

func (e *ElasticSearchClient) SendMessage(message kafka.Message) error {
	var messageObj config.Message
	json.Unmarshal(message.Value, &messageObj)
	messageObj.Time = message.Time.Nanosecond()

	key := fmt.Sprintf("%v_%v_%v", message.Topic, message.Partition, message.Offset)
	index := fmt.Sprintf("clan_%v", messageObj.ClanId)
	log.Println(messageObj, key, index)
	req := esapi.IndexRequest{
		Index:        index,
		Body:         strings.NewReader(string(message.Value)),
		DocumentID:   key,
		DocumentType: "message",
		Refresh:      "true",
	}
	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return err
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

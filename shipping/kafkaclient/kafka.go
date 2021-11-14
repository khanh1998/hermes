package kafkaclient

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/shipping/config"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	conn   *kafka.Conn
	reader *kafka.Reader
}

func NewKafkaClient(env *config.Env) (*KafkaClient, error) {
	conn, err := kafka.DialLeader(context.Background(), env.KAFKA_NETWORK_PROTOCOL, env.KAFKA_URI, env.KAFKA_TOPIC, 0)
	if err != nil {
		return nil, err
	}
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{env.KAFKA_URI},
		Topic:          env.KAFKA_TOPIC,
		Partition:      0,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		GroupID:        "shippers",
	})
	return &KafkaClient{
		conn:   conn,
		reader: r,
	}, nil
}

func (k *KafkaClient) ReceiveMessage() ([]byte, error) {
	log.Printf("lag: %v offset: %v", k.reader.Lag(), k.reader.Offset())
	m, err := k.reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	var message config.Message
	json.Unmarshal(m.Value, &message)
	fmt.Printf("message at offset %d: %s = %v\n", m.Offset, string(m.Key), message)
	return m.Value, nil
}

func (k *KafkaClient) Close() error {
	return k.conn.Close()
}

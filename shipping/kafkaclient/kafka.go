package kafkaclient

import (
	"context"
	"hermes/shipping/config"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	conn *kafka.Conn
}

func NewKafkaClient(env *config.Env) (*KafkaClient, error) {
	conn, err := kafka.DialLeader(context.Background(), env.KAFKA_NETWORK_PROTOCOL, env.KAFKA_URI, env.KAFKA_TOPIC, 0)
	if err != nil {
		return nil, err
	}
	return &KafkaClient{
		conn: conn,
	}, nil
}

func (k *KafkaClient) ReceiveMessage() ([]byte, error) {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "message",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)
	m, err := r.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	// var message config.Message
	// json.Unmarshal(m.Value, &message)
	return m.Value, nil
}

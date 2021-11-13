package kafkaclient

import (
	"context"
	"encoding/json"
	"hermes/socket/config"
	"time"

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

func (k *KafkaClient) SendMessage(message *config.Message) error {
	k.conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	m, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = k.conn.Write(m)
	if err != nil {
		return err
	}
	return nil
}

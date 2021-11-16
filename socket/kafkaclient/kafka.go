package kafkaclient

import (
	"context"
	"encoding/json"
	"hermes/socket/config"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	writer *kafka.Writer
}

func NewKafkaClient(env *config.Env) (*KafkaClient, error) {
	writer := kafka.Writer{
		Addr:     kafka.TCP(env.KAFKA_URI),
		Topic:    env.KAFKA_TOPIC,
		Balancer: &kafka.LeastBytes{},
	}
	return &KafkaClient{
		writer: &writer,
	}, nil
}

func (k *KafkaClient) SendMessage(message *config.Message) error {
	m, err := json.Marshal(message)
	if err != nil {
		return err
	}
	k.writer.WriteMessages(context.Background(), kafka.Message{
		Value: m,
	})
	if err != nil {
		return err
	}
	return nil
}

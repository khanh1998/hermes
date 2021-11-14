package main

import (
	"hermes/shipping/config"
	"hermes/shipping/esclient"
	"hermes/shipping/kafkaclient"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env, err := config.GetEnv()
	if err != nil {
		log.Println(err)
	}

	es, err := esclient.NewElasticSearchClient(env)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connect to elasticsearch")

	kafka, err := kafkaclient.NewKafkaClient(env)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connect to kafka")

	// catch interupt event to close kafka connection gently
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt, syscall.SIGTERM)
	go func() {
		_, ok := <-c
		if ok {
			log.Println("close kafka connection")
			if err := kafka.Close(); err != nil {
				log.Println("close kafka connection fail")
			}
			log.Println("close kafka successfully")
			os.Exit(0)
		}
	}()

	for {
		// fetchs message from Kafka
		m, err := kafka.ReceiveMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		// pushs message to Elasticsearch
		if err := es.SendMessage(m); err != nil {
			log.Println(err)
			continue
		}
		// confirms to Kafka that the message is delivered successfully,
		// and the Kafka isn't gonna send this message again.
		if err := kafka.CommitMessage(m); err != nil {
			log.Println(err)
			continue
		}
	}
}

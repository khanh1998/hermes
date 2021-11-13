package main

import (
	"hermes/shipping/config"
	"hermes/shipping/esclient"
	"hermes/shipping/kafkaclient"
	"log"
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

	kafka, err := kafkaclient.NewKafkaClient(env)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		m, err := kafka.ReceiveMessage()
		if err != nil {
			log.Println(err)
			break
		}
		err = es.SendMessage(m)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

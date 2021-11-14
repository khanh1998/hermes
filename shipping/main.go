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

	kafka, err := kafkaclient.NewKafkaClient(env)
	if err != nil {
		log.Println(err)
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt, syscall.SIGTERM)
	go func() {
		_, ok := <-c
		if ok {
			log.Println("close kafka connection")
			if err := kafka.Close(); err != nil {
				log.Println("close connection fail")
			}
			os.Exit(0)
		}
	}()

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

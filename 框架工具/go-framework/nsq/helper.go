package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

const (
	NSQD_ADDR  = "127.0.0.1:4150"
	TOPIC_TEST = "test"
)

func newProducer(addr string) *nsq.Producer {
	var (
		cfg           = nsq.NewConfig()
		producer, err = nsq.NewProducer(addr, cfg)
	)

	if err != nil {
		log.Fatal(err)
	}

	return producer
}

func newConsumer(topic, channel string) *nsq.Consumer {
	var (
		cfg           = nsq.NewConfig()
		consumer, err = nsq.NewConsumer(topic, channel, cfg)
	)

	if err != nil {
		log.Fatal(err)
	}

	return consumer
}

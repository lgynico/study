package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestSub(t *testing.T) {
	nc := connect(nats.DefaultURL)
	defer nc.Close()

	_, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	})

	if err != nil {
		log.Fatal(err)
	}

	blocking()
}

func TestSubChan(t *testing.T) {
	nc := connect(nats.DefaultURL)
	defer nc.Close()

	ch := make(chan *nats.Msg, 64)
	if _, err := nc.ChanSubscribe("foo", ch); err != nil {
		log.Fatal(err)
	}

	msg := <-ch

	fmt.Printf("Received a message: %s\n", string(msg.Data))
	fmt.Println("Subscriber exit!!")
}

func TestPub(t *testing.T) {
	nc := connect(nats.DefaultURL)
	defer nc.Close()

	if err := nc.Publish("foo", []byte("hello world")); err != nil {
		log.Fatal(err)
	}
}

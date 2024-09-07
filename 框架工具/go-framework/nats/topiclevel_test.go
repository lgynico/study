package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func TestTopicLevel(t *testing.T) {
	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(3)

	go func() {
		defer wg.Done()

		nc := connect(nats.DefaultURL)
		defer nc.Close()

		ch := make(chan *nats.Msg, 64)
		_, err := nc.ChanSubscribe("player.*.login", ch)
		if err != nil {
			log.Fatal(err)
		}

		timeout := time.NewTimer(15 * time.Second)
	LOOP:
		for {
			select {
			case msg := <-ch:
				fmt.Printf("player login: %s\n", string(msg.Data))
			case <-timeout.C:
				break LOOP
			}
		}

	}()

	go func() {
		defer wg.Done()

		nc := connect(nats.DefaultURL)
		defer nc.Close()

		ch := make(chan *nats.Msg, 64)
		_, err := nc.ChanSubscribe("player.1.*", ch)
		if err != nil {
			log.Fatal(err)
		}

		timeout := time.NewTimer(15 * time.Second)
	LOOP:
		for {
			select {
			case msg := <-ch:
				fmt.Printf("player 1: %s\n", string(msg.Data))
			case <-timeout.C:
				break LOOP
			}
		}

	}()

	go func() {
		defer wg.Done()

		nc := connect(nats.DefaultURL)
		defer nc.Close()

		ch := make(chan *nats.Msg, 64)
		_, err := nc.ChanSubscribe("player.>", ch)
		if err != nil {
			log.Fatal(err)
		}

		timeout := time.NewTimer(15 * time.Second)
	LOOP:
		for {
			select {
			case msg := <-ch:
				fmt.Printf("player event: %s\n", string(msg.Data))
			case <-timeout.C:
				break LOOP
			}
		}

	}()

	go pubTopic()
}

func pubTopic() {
	nc := connect(nats.DefaultURL)
	defer nc.Close()

	for i := 1; i <= 10; i++ {
		if err := nc.Publish(fmt.Sprintf("player.%d.login", i), []byte(strconv.Itoa(i))); err != nil {
			log.Fatal(err)
		}

		if err := nc.Publish(fmt.Sprintf("player.%d.enter", i), []byte("enter")); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}
}

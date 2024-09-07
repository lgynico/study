package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func TestRequest(t *testing.T) {
	nc := connect(nats.DefaultURL)
	defer nc.Close()

	for i := 0; i < 10; i++ {
		msg, err := nc.Request("foo", []byte("help me"), 3*time.Second)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Receive response: %s\n", string(msg.Data))

		time.Sleep(time.Second)
	}

}

func TestRespond(t *testing.T) {
	var (
		nSubscriber = 2
		wg          sync.WaitGroup
	)

	defer wg.Wait()

	for i := 0; i < nSubscriber; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			nc := connect(nats.DefaultURL)
			defer nc.Close()

			if _, err := nc.Subscribe("foo", func(msg *nats.Msg) {
				fmt.Printf("Answer_%d receive message: %s\n", i, string(msg.Data))
				msg.Respond([]byte(fmt.Sprintf("This is answer_%d comming...", i)))
			}); err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Minute)
		}(i)

	}
}

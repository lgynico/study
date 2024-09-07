package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
)

const (
	fetchStreamName = "fetchStream"
	fetchSubjectAll = "fetchSubject.*"
	fetchSubject    = "fetchSubject.key1"
)

func TestProducer(t *testing.T) {
	nc := connect(nats.DefaultURL)
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	ctx, f := context.WithTimeout(context.Background(), 60*time.Second)
	defer f()

	info, err := js.StreamInfo(fetchStreamName)
	// if err != nats.ErrStreamNotFound {
	// 	log.Fatal(err)
	// }

	if info == nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:       fetchStreamName,
			Subjects:   []string{fetchSubjectAll},
			Retention:  nats.WorkQueuePolicy,
			Replicas:   1,
			Discard:    nats.DiscardOld,
			Duplicates: 30 * time.Second,
		}, nats.Context(ctx))
		if err != nil {
			log.Fatalf("can't add: %v\n", err)
		}
	}

	var (
		results       = make(chan int64)
		totalTime     int64
		totalMessages int64
	)

	go func() {
		i := 0
		for {
			js.Publish(fetchSubject, []byte(fmt.Sprintf("message==%d", i)), nats.Context(ctx))
			log.Printf("[publisher] sent %d\n", i)
			time.Sleep(time.Second)
			i++
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("sent %d messages with average time of %f\n", totalMessages, math.Round(float64(totalTime/totalMessages)))
			js.DeleteStream(fetchStreamName)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

func TestConsumer(t *testing.T) {
	ctx, fun := context.WithTimeout(context.Background(), 60*time.Second)
	defer fun()

	id := uuid.NewV4().String()
	nc := connect(nats.DefaultURL, nats.Name(id))
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	sub, _ := js.PullSubscribe(fetchSubject, "group")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for {
		msgs, err := sub.Fetch(1, nats.Context(ctx))
		if err != nil {
			log.Fatal(err)
		}

		msg := msgs[0]
		log.Printf("[Consumer %s] received msg (%v)\n", id, string(msg.Data))
		if err = msg.Ack(nats.Context(ctx)); err != nil {
			log.Fatal(err)
		}
	}
}

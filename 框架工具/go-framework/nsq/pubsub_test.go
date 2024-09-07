package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
)

func TestPub(t *testing.T) {
	var (
		count    int
		producer = newProducer(NSQD_ADDR)
	)

	for {
		count++
		msg := fmt.Sprintf("test %d", count)
		fmt.Println("发布消息：", msg)
		if err := producer.Publish(TOPIC_TEST, []byte(msg)); err != nil {
			log.Fatal("publish error: " + err.Error())
		}

		time.Sleep(time.Second)
	}
}

// LoadBalance: same topic and channel
func TestLoadBalance(t *testing.T) {
	var (
		nConsumer = 2
		wg        sync.WaitGroup
	)

	defer wg.Wait()

	for i := 0; i < nConsumer; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			consumer := newConsumer(TOPIC_TEST, "channel0")

			consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
				log.Printf("消费者 [%d] 收到消息：%s\n", i, string(message.Body))
				return nil
			}))

			if err := consumer.ConnectToNSQD(NSQD_ADDR); err != nil {
				log.Fatal(err)
			}

			<-consumer.StopChan
		}(i)
	}

}

// Broadcast: same topic, different channel
func TestSubBroadcast(t *testing.T) {
	var (
		nConsumer = 3
		wg        sync.WaitGroup
	)

	defer wg.Wait()

	for i := 0; i < nConsumer; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			consumer := newConsumer(TOPIC_TEST, fmt.Sprintf("channel%d", i))

			consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
				log.Printf("消费者 [%d] 收到消费：%s\n", i, string(message.Body))
				return nil
			}))

			if err := consumer.ConnectToNSQD(NSQD_ADDR); err != nil {
				log.Fatal(err)
			}

			<-consumer.StopChan
		}(i)
	}
}

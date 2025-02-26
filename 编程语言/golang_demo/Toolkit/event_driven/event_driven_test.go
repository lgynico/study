package eventdriven

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channelEvent()
}

func TestEventBus(t *testing.T) {
	bus := &EventBus{}
	bus.Subscribe("UserRegistered", func(data any) {
		fmt.Println("Handle user registered event:", data)
	})

	bus.Publish("UserRegistered", "Nico")
}

func TestKafka(t *testing.T) {
	go kafkaProducer()
	go kafkaComsumer()

	time.Sleep(time.Second * 5)
}

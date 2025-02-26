package eventdriven

import (
	"fmt"
	"time"
)

type Event struct {
	Name string
	Data any
}

// 基于 channel 的事件驱动
func channelEvent() {
	eventC := make(chan Event, 10)

	go func() {
		for event := range eventC {
			fmt.Printf("[Consumer] Received event: %+v\n", event)
		}
	}()

	eventC <- Event{Name: "UserRegistered", Data: "Nico"}
	time.Sleep(time.Second)
}

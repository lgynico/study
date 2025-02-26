package pattern

import (
	"context"
	"fmt"
	"time"
)

func cancelByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout after 5 seconds")
	case <-ctx.Done():
		fmt.Println("canceled by context")
	}
}

package rands

import (
	"fmt"
	"testing"
)

func TestRandomEven(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandomEven(1, 100))
	}
}

package bit

import (
	"fmt"
	"testing"
)

// 计算 n 里面有几个 1
func countOne(n int) int {
	var count int
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func TestCountOne(t *testing.T) {
	fmt.Println(countOne(0b11111))
	fmt.Println(countOne(0))
	fmt.Println(countOne(1))
	fmt.Println(countOne(0b101101110111011))
}

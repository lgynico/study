package heap

import (
	"fmt"
	"testing"
)

func TestDefault(t *testing.T) {
	heap := Default[int, string]{}
	heap.Init()

	heap.Put(3, "Nico")
	heap.Put(1, "Mike")
	heap.Put(2, "Jeck")
	heap.Put(4, "Amy")
	heap.Put(5, "Harry")

	key, value, _ := heap.Min()
	fmt.Println(key, value)

	key, value, _ = heap.Min()
	fmt.Println(key, value)

	key, value, _ = heap.Min()
	fmt.Println(key, value)

	key, value, _ = heap.Min()
	fmt.Println(key, value)

	key, value, _ = heap.Min()
	fmt.Println(key, value)

	key, value, _ = heap.Min()
	fmt.Println(key, value)
}

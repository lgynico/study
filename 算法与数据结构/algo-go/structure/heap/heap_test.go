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

func TestArray(t *testing.T) {
	heap := NewArray(func(a, b int) int { return a - b })
	heap.Push(5)
	heap.Push(3)
	fmt.Println(heap.Pop())
	heap.Push(7)
	heap.Push(3)
	heap.Push(2)
	heap.Push(9)
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
}

func TestAdvance(t *testing.T) {
	heap := NewAdvance(func(a, b int) int { return a - b })
	heap.Push(5)
	heap.Push(3)
	fmt.Println(heap.Pop())
	heap.Push(7)
	heap.Push(3)
	heap.Push(2)
	heap.Push(9)
	heap.Push(11)
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())

	heap.Push(1)
	heap.Push(3)
	heap.Push(2)
	heap.Push(4)
	heap.Remove(2)
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())

}

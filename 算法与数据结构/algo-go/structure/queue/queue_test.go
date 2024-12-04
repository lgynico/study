package queue

import (
	"fmt"
	"testing"
)

func TestQueueList(t *testing.T) {
	queue := List[int]{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	fmt.Println("queue len:", queue.Size())

	var e int
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
}

func TestQueueArray(t *testing.T) {
	queue := Array[int]{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	fmt.Println("queue len:", queue.Size())

	var e int
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
}

func TestDeque(t *testing.T) {
	deque := Deque[int]{}

	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)

	fmt.Println("deque len:", deque.Size())

	e, _ := deque.Dequeue()
	fmt.Println(e)
	e, _ = deque.Dequeue()
	fmt.Println(e)
	e, _ = deque.Dequeue()
	fmt.Println(e)
	e, _ = deque.Dequeue()
	fmt.Println(e)
	e, _ = deque.Dequeue()
	fmt.Println(e)

	fmt.Println("deque is empty:", deque.IsEmpty())

	deque.EnqueueFront(5)
	deque.EnqueueFront(4)
	deque.EnqueueFront(3)
	deque.EnqueueFront(2)
	deque.EnqueueFront(1)

	fmt.Println("deque len:", deque.Size())

	e, _ = deque.DequeueLast()
	fmt.Println(e)
	e, _ = deque.DequeueLast()
	fmt.Println(e)
	e, _ = deque.DequeueLast()
	fmt.Println(e)
	e, _ = deque.DequeueLast()
	fmt.Println(e)
	e, _ = deque.DequeueLast()
	fmt.Println(e)

	fmt.Println("deque len:", deque.Size())
}

package misc

import (
	"fmt"
	"testing"
)

func TestQueueStack(t *testing.T) {
	stack := QueueStack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	fmt.Println("Stack len:", stack.Size())
	e, _ := stack.Peek()
	fmt.Println("Stack top:", e)

	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)
	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)
	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)

	stack.Push(6)

	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)
	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)

	fmt.Println("Stack is empty:", stack.IsEmpty())

	_, ok := stack.Pop()
	fmt.Println("Stack top:", ok)
}

func TestQueueArray(t *testing.T) {
	queue := StackQueue[int]{}
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

	queue.Enqueue(6)

	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
	e, _ = queue.Dequeue()
	fmt.Println(e)
}

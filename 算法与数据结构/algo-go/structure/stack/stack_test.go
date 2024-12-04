package stack

import (
	"fmt"
	"testing"
)

func TestStackList(t *testing.T) {
	stack := List[int]{}

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
	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)
	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)

	fmt.Println("Stack is empty:", stack.IsEmpty())
}

func TestStackArray(t *testing.T) {
	stack := Array[int]{}

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
	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)
	e, _ = stack.Pop()
	fmt.Println("Stack top:", e)

	fmt.Println("Stack is empty:", stack.IsEmpty())

	_, ok := stack.Pop()
	fmt.Println("Stack top:", ok)
}

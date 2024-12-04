package misc

import "github.com/lgynico/algo-go/structure/stack"

// 使用栈实现单向队列
type StackQueue[T any] struct {
	enqueue, dequeue stack.Array[T]
}

func (p *StackQueue[T]) Enqueue(element T) {
	p.enqueue.Push(element)
}

func (p *StackQueue[T]) Dequeue() (element T, ok bool) {
	p.checkDequeue()

	if !p.dequeue.IsEmpty() {
		element, ok = p.dequeue.Pop()
	}

	return
}

func (p *StackQueue[T]) Peek() (element T, ok bool) {
	p.checkDequeue()

	if !p.dequeue.IsEmpty() {
		element, ok = p.dequeue.Peek()
	}

	return
}

func (p *StackQueue[T]) IsEmpty() bool {
	return p.dequeue.IsEmpty() && p.enqueue.IsEmpty()
}

func (p *StackQueue[T]) Size() int {
	return p.dequeue.Size() + p.enqueue.Size()
}

func (p *StackQueue[T]) checkDequeue() {
	if p.dequeue.IsEmpty() {
		for !p.enqueue.IsEmpty() {
			e, _ := p.enqueue.Pop()
			p.dequeue.Push(e)
		}
	}
}

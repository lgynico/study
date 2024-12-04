package misc

import (
	"github.com/lgynico/algo-go/structure/queue"
)

// 使用单向队列实现栈
type QueueStack[T any] struct {
	push, pop queue.Array[T]
}

func (p *QueueStack[T]) Push(element T) {
	p.push.Enqueue(element)
}

func (p *QueueStack[T]) Pop() (element T, ok bool) {
	p.pushToPop()

	if !p.push.IsEmpty() {
		element, ok = p.push.Dequeue()
	}

	p.push, p.pop = p.pop, p.push

	return
}

func (p *QueueStack[T]) Peek() (element T, ok bool) {
	p.pushToPop()
	if !p.push.IsEmpty() {
		element, ok = p.push.Peek()
	}
	return
}

func (p *QueueStack[T]) Size() int {
	return p.push.Size() + p.pop.Size()
}

func (p *QueueStack[T]) IsEmpty() bool {
	return p.Size() == 0
}

func (p *QueueStack[T]) pushToPop() {
	for p.push.Size() > 1 {
		e, _ := p.push.Dequeue()
		p.pop.Enqueue(e)
	}
}

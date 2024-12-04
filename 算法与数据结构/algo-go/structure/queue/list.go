package queue

import "github.com/lgynico/algo-go/structure/list"

type List[T any] struct {
	elements list.Array[T]
}

func (p *List[T]) Enqueue(element T) {
	p.elements.Add(element)
}

func (p *List[T]) Dequeue() (element T, ok bool) {
	element, ok = p.elements.RemoveFront()
	return
}

func (p *List[T]) Peek() (element T, ok bool) {
	element, ok = p.elements.Get(0)
	return
}

func (p *List[T]) IsEmpty() bool { return p.elements.IsEmpty() }

func (p *List[T]) Size() int {
	return p.elements.Size()
}

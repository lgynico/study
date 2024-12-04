package stack

import "github.com/lgynico/algo-go/structure/list"

type List[T any] struct {
	elements list.Array[T]
}

func (p *List[T]) Push(element T) {
	p.elements.Add(element)
}

func (p *List[T]) Pop() (element T, ok bool) {
	element, ok = p.elements.RemoveLast()
	return
}

func (p *List[T]) Peek() (element T, ok bool) {
	element, ok = p.elements.GetLast()
	return
}

func (p *List[T]) Size() int {
	return p.elements.Size()
}

func (p *List[T]) IsEmpty() bool {
	return p.elements.IsEmpty()
}

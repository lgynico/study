package heap

import (
	"github.com/lgynico/algo-go/utils/constraints"
)

type heapImpl[T constraints.Comparable] struct {
	elements []T
	size     int
}

func (p *heapImpl[T]) Push(x any) {
	p.elements = append(p.elements, x.(T))
	p.size++
}

func (p *heapImpl[T]) Pop() any {
	e := p.elements[p.size-1]
	p.size--
	// p.elements = append(p.elements[:0], p.elements[1:]...)

	return e
}

func (p *heapImpl[T]) Len() int {
	return p.size
}

func (p *heapImpl[T]) Less(i, j int) bool {
	return p.elements[i] < p.elements[j]
}

func (p *heapImpl[T]) Swap(i, j int) {
	p.elements[i], p.elements[j] = p.elements[j], p.elements[i]
}

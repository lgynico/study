package queue

type Array[T any] struct {
	elements []T
}

func (p *Array[T]) Enqueue(element T) {
	p.elements = append(p.elements, element)
}

func (p *Array[T]) Dequeue() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[0]
		ok = true

		p.elements = append(p.elements[:0], p.elements[1:]...)
	}
	return
}

func (p *Array[T]) Peek() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[0]
		ok = true
	}

	return
}

func (p *Array[T]) IsEmpty() bool {
	return len(p.elements) == 0
}

func (p *Array[T]) Size() int {
	return len(p.elements)
}

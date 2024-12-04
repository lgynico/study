package queue

type Deque[T any] struct {
	elements []T
	size     int
}

func (p *Deque[T]) Enqueue(element T) {
	p.elements = append(p.elements, element)
	p.size++
}

func (p *Deque[T]) Dequeue() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[0]
		ok = true

		p.elements = append(p.elements[:0], p.elements[1:]...)
		p.size--
	}

	return
}

func (p *Deque[T]) EnqueueFront(element T) {
	p.elements = append([]T{element}, p.elements...)
	p.size++
}

func (p *Deque[T]) DequeueLast() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[p.Size()-1]
		ok = true

		p.elements = append(p.elements[:p.Size()-1], p.elements[:p.Size()]...)
		p.size--
	}

	return
}

func (p *Deque[T]) PeekFront() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[0]
		ok = true
	}

	return
}

func (p *Deque[T]) PeekLast() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[p.Size()-1]
		ok = true
	}

	return
}

func (p *Deque[T]) Size() int {
	return p.size
}

func (p *Deque[T]) IsEmpty() bool {
	return p.size == 0
}

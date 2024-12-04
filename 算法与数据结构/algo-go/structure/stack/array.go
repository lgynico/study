package stack

type Array[T any] struct {
	elements []T
}

func (p *Array[T]) Push(element T) {
	p.elements = append(p.elements, element)
}

func (p *Array[T]) Pop() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[p.Size()-1]
		ok = true

		p.elements = append(p.elements[:p.Size()-1], p.elements[p.Size():]...)
	}
	return
}

func (p *Array[T]) Peek() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[p.Size()-1]
		ok = true
	}
	return
}

func (p *Array[T]) Size() int { return len(p.elements) }

func (p *Array[T]) IsEmpty() bool { return len(p.elements) == 0 }

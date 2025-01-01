package heap

type (
	HeapComparator[T any] func(a, b T) int

	Array[T any] struct {
		elements []T
		size     int
		compare  HeapComparator[T]
	}
)

func NewArray[T any](compare HeapComparator[T]) Array[T] {
	return Array[T]{
		elements: make([]T, 0),
		compare:  compare,
	}
}

func (p *Array[T]) Push(element T) {
	if p.size < len(p.elements) {
		p.elements[p.size] = element
	} else {
		p.elements = append(p.elements, element)
	}
	p.size++

	p.heapInsert(p.size - 1)
}

func (p *Array[T]) Pop() T {
	element := p.elements[0]
	p.elements[0], p.elements[p.size-1] = p.elements[p.size-1], p.elements[0]
	p.size--

	p.heapify(0)

	return element
}

func (p *Array[T]) Peek() T {
	return p.elements[0]
}

func (p *Array[T]) Size() int {
	return p.size
}

func (p *Array[T]) IsEmpty() bool {
	return p.Size() == 0
}

func (p *Array[T]) heapInsert(i int) {
	for i > 0 && p.compare(p.elements[i], p.elements[(i-1)/2]) < 0 {
		p.elements[i], p.elements[(i-1)/2] = p.elements[(i-1)/2], p.elements[i]
		i = (i - 1) / 2
	}
}

func (p *Array[T]) heapify(i int) {
	for i < p.size {
		child := 2*i + 1
		if child >= p.size {
			break
		}
		if child+1 < p.size && p.compare(p.elements[child], p.elements[child+1]) > 0 {
			child++
		}
		if p.compare(p.elements[i], p.elements[child]) <= 0 {
			break
		}
		p.elements[i], p.elements[child] = p.elements[child], p.elements[i]
		i = child
	}
}

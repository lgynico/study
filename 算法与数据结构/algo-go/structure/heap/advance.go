package heap

type Advance[T any] struct {
	elements []T
	size     int
	indexes  map[any]int
	compare  HeapComparator[T]
}

func NewAdvance[T any](comprator HeapComparator[T]) Advance[T] {
	return Advance[T]{
		elements: make([]T, 0),
		size:     0,
		indexes:  make(map[any]int),
		compare:  comprator,
	}
}

func (p *Advance[T]) Push(element T) {
	if p.size == len(p.elements) {
		p.elements = append(p.elements, element)
	} else {
		p.elements[p.size] = element
	}
	p.indexes[element] = p.size
	p.heapInsert(p.size)
	p.size++
}

func (p *Advance[T]) Pop() T {
	element := p.elements[0]
	p.elements[0] = p.elements[p.size-1]
	p.size--
	delete(p.indexes, element)

	if p.size > 0 {
		p.indexes[p.elements[0]] = 0
		p.heapify(0, p.size-1)
	}

	return element
}

func (p *Advance[T]) Peek() T {
	return p.elements[0]
}

func (p *Advance[T]) Remove(element T) {
	i, ok := p.indexes[element]
	if !ok {
		return
	}

	delete(p.indexes, element)
	p.elements[i] = p.elements[p.size-1]
	p.size--
	if i == p.size {
		return
	}

	p.indexes[p.elements[i]] = i
	p.heapify(i, p.size-1)
}

func (p *Advance[T]) Contains(element T) bool {
	_, ok := p.indexes[element]
	return ok
}

func (p *Advance[T]) Size() int {
	return p.size
}

func (p *Advance[T]) IsEmpty() bool {
	return p.size == 0
}

func (p *Advance[T]) Elements() []T {
	return p.elements[:p.size]
}

func (p *Advance[T]) heapInsert(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if p.compare(p.elements[i], p.elements[parent]) >= 0 {
			break
		}

		p.swap(i, parent)
		i = parent
	}
}

func (p *Advance[T]) heapify(i, j int) {
	for i <= j {
		child := 2*i + 1
		if child > j {
			break
		}
		if child+1 <= j && p.compare(p.elements[child], p.elements[child+1]) > 0 {
			child++
		}
		if p.compare(p.elements[i], p.elements[child]) <= 0 {
			break
		}

		p.swap(i, child)
		i = child
	}
}

func (p *Advance[T]) Resign(element T) {
	p.heapInsert(p.indexes[element])
	p.heapify(p.indexes[element], p.size-1)
}

func (p *Advance[T]) swap(i, j int) {
	p.elements[i], p.elements[j] = p.elements[j], p.elements[i]
	p.indexes[p.elements[i]] = i
	p.indexes[p.elements[j]] = j
}

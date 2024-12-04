package list

import (
	"fmt"
	"strings"
)

type Array[T any] struct {
	elements []T
}

func (p *Array[T]) Get(index int) (element T, ok bool) {
	if index < 0 || index >= p.Size() {
		return
	}

	element = p.elements[index]
	ok = true

	return
}

func (p *Array[T]) GetLast() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements[p.Size()-1]
		ok = true
	}

	return
}

func (p *Array[T]) Add(element T) {
	p.elements = append(p.elements, element)
}

func (p *Array[T]) AddFront(element T) {
	p.elements = append([]T{element}, p.elements...)
}

func (p *Array[T]) Remove(element T) {
	var (
		size  = p.Size()
		index = 0
	)

	for index < size {
		if any(p.elements[index]) != any(element) {
			index++
			continue
		}

		p.elements = append(p.elements[:index], p.elements[index+1:]...)
		size--
	}
}

func (p *Array[T]) RemoveFront() (element T, ok bool) {
	return p.RemoveAt(0)
}

func (p *Array[T]) RemoveLast() (element T, ok bool) {
	return p.RemoveAt(p.Size() - 1)
}

func (p *Array[T]) RemoveAt(index int) (element T, ok bool) {
	if index < 0 || p.Size() <= index {
		return
	}

	element = p.elements[index]
	ok = true

	p.elements = append(p.elements[:index], p.elements[index+1:]...)

	return
}

func (p *Array[T]) Reverse() {
	if p.Size() <= 1 {
		return
	}

	for i := 0; i < p.Size()>>1; i++ {
		p.elements[i], p.elements[p.Size()-1-i] = p.elements[p.Size()-1-i], p.elements[i]
	}
}

func (p *Array[T]) Size() int {
	return len(p.elements)
}

func (p *Array[T]) IsEmpty() bool {
	return len(p.elements) == 0
}

func (p *Array[T]) String() string {
	if p.IsEmpty() {
		return ""
	}

	sb := strings.Builder{}
	for _, element := range p.elements {
		sb.WriteString(fmt.Sprintf(" -> %v", element))
	}
	return sb.String()
}

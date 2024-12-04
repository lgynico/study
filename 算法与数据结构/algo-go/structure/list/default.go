package list

import (
	"container/list"
	"fmt"
	"strings"
)

type Default[T any] struct {
	elements list.List
}

func (p *Default[T]) Get(index int) (element T, ok bool) {
	if index < 0 || index >= p.Size() {
		return
	}

	node := p.elements.Front()
	for i := 0; i < index; i++ {
		node = node.Next()
	}

	element = node.Value.(T)
	ok = true

	return
}

func (p *Default[T]) GetLast() (element T, ok bool) {
	if !p.IsEmpty() {
		element = p.elements.Back().Value.(T)
		ok = true
	}

	return
}

func (p *Default[T]) Add(element T) {
	p.elements.PushBack(element)
}

func (p *Default[T]) AddFront(element T) {
	p.elements.PushFront(element)
}

func (p *Default[T]) Remove(element T) {
	for node := p.elements.Front(); node != nil; node = node.Next() {
		if node.Value == any(element) {
			p.elements.Remove(node)
		}
	}
}

func (p *Default[T]) RemoveFront() (element T, ok bool) {
	if !p.IsEmpty() {
		node := p.elements.Front()
		p.elements.Remove(node)

		element = node.Value.(T)
		ok = true
	}
	return
}

func (p *Default[T]) RemoveLast() (element T, ok bool) {
	if !p.IsEmpty() {
		node := p.elements.Back()
		p.elements.Remove(node)

		element = node.Value.(T)
		ok = true
	}

	return
}

func (p *Default[T]) RemoveAt(index int) (element T, ok bool) {
	if index < 0 || index >= p.Size() {
		return
	}

	node := p.elements.Front()
	for i := 1; i < index; i++ {
		node = node.Next()
	}

	p.elements.Remove(node)

	element = node.Value.(T)
	ok = true

	return
}

func (p *Default[T]) Reverse() {
	if p.Size() <= 1 {
		return
	}

	head := p.elements.Front()

	for {
		node := p.elements.Back()
		if node == head {
			break
		}

		p.elements.MoveBefore(node, head)
	}
}

func (p *Default[T]) Size() int {
	return p.elements.Len()
}

func (p *Default[T]) IsEmpty() bool {
	return p.elements.Len() == 0
}

func (p *Default[T]) String() string {
	if p.IsEmpty() {
		return ""
	}

	sb := strings.Builder{}
	for node := p.elements.Front(); node != nil; node = node.Next() {
		sb.WriteString(fmt.Sprintf("-> %v", node.Value))
	}
	sb.WriteString("\r\n")

	for node := p.elements.Back(); node != nil; node = node.Prev() {
		sb.WriteString(fmt.Sprintf("<- %v", node.Value))
	}

	return sb.String()
}

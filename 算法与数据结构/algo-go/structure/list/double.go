package list

import (
	"fmt"
	"strings"
)

type (
	dnode[T any] struct {
		data       T
		prev, next *dnode[T]
	}

	Double[T any] struct {
		head, tail *dnode[T]
		size       int
	}
)

func (p *Double[T]) Get(index int) (element T, ok bool) {
	if p.IsEmpty() || p.size <= index {
		return
	}

	node := p.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	element = node.data
	ok = true

	return
}

func (p *Double[T]) Add(element T) {
	node := dnode[T]{data: element}
	if p.IsEmpty() {
		p.head = &node
		p.tail = p.head
	} else {
		p.tail.next = &node
		node.prev = p.tail
		p.tail = &node
	}

	p.size++
}

func (p *Double[T]) AddFront(element T) {
	node := dnode[T]{data: element}
	node.next = p.head
	p.head = &node
	if node.next != nil {
		node.next.prev = &node
	}
	p.size++
}

func (p *Double[T]) Remove(element T) {
	if p.IsEmpty() {
		return
	}

	var (
		dummy = dnode[T]{next: p.head}
		curr  = p.head
	)

	curr.prev = &dummy

	for curr != nil {
		if any(curr.data) == any(element) {
			curr.prev.next = curr.next
			if curr.next != nil {
				curr.next.prev = curr.prev
			}
			if curr == p.tail {
				p.tail = curr.prev
			}
			p.size--
		}
		curr = curr.next
	}

	p.head = dummy.next
	if p.head != nil {
		p.head.prev = nil
	}

	if p.tail == &dummy {
		p.tail = nil
	}

}

func (p *Double[T]) RemoveFront() (element T, ok bool) {
	if p.IsEmpty() {
		return
	}

	element = p.head.data
	ok = true

	p.head = p.head.next
	if p.head != nil {
		p.head.prev = nil
	} else {
		p.tail = nil
	}

	p.size--

	return
}

func (p *Double[T]) RemoveAt(index int) (element T, ok bool) {
	if p.IsEmpty() || p.Size() <= index {
		return
	}

	node := p.head

	for i := 1; i < index; i++ {
		node = node.next
	}

	if node == p.head {
		p.head = node.next
		p.head.prev = nil
	} else {
		node.prev.next = node.next
		if node.next != nil {
			node.next.prev = node.prev
		}
	}

	if node == p.tail {
		p.tail = node.prev
	}

	p.size--

	element = node.data
	ok = true

	return
}

func (p *Double[T]) Reverse() {
	if p.Size() <= 1 {
		return
	}

	curr := p.head

	for curr != nil {
		next := curr.next
		curr.next = curr.prev
		curr.prev = next
		curr = next
	}

	p.head, p.tail = p.tail, p.head
}

func (p *Double[T]) Size() int {
	return p.size
}

func (p *Double[T]) IsEmpty() bool {
	return p.size == 0
}

func (p *Double[T]) String() string {
	if p.IsEmpty() {
		return ""
	}

	sb := strings.Builder{}
	curr := p.head
	for curr != nil {
		sb.WriteString(fmt.Sprintf("-> %v", curr.data))
		curr = curr.next
	}
	sb.WriteString("\r\n")

	curr = p.tail
	for curr != nil {
		sb.WriteString(fmt.Sprintf("<- %v", curr.data))
		curr = curr.prev
	}

	return sb.String()
}

package list

import (
	"fmt"
	"strings"
)

type (
	snode[T any] struct {
		data T
		next *snode[T]
	}

	Single[T any] struct {
		head, tail *snode[T]
		size       int
	}
)

func (p *Single[T]) Get(index int) (element T, ok bool) {
	if index <= p.size {
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

func (p *Single[T]) Add(element T) {
	node := snode[T]{data: element}
	if p.IsEmpty() {
		p.head = &node
		p.tail = &node
	} else {
		p.tail.next = &node
		p.tail = &node
	}

	p.size++
}

func (p *Single[T]) Remove(element T) {
	if p.IsEmpty() {
		return
	}

	p.head = p.remove(element)
}

func (p *Single[T]) RemoveAt(index int) (element T, ok bool) {
	if p.IsEmpty() || p.Size() <= index {
		return
	}

	if index == 0 {
		if p.tail == p.head {
			p.tail = p.head.next
		}
		p.head = p.head.next
	}

	var (
		prev = p.head
		node = prev.next
	)
	for i := 1; i < index; i++ {
		prev = node
		node = node.next
	}

	prev.next = node.next
	if p.tail == node {
		p.tail = prev
	}

	p.size--

	element = node.data
	ok = true

	return
}

func (p *Single[T]) Reverse() {
	if p.Size() <= 1 {
		return
	}

	var (
		prev *snode[T]
		curr = p.head
	)

	p.head, p.tail = p.tail, p.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
}

func (p *Single[T]) Size() int {
	return p.size
}

func (p *Single[T]) IsEmpty() bool {
	return p.size == 0
}

func (p *Single[T]) String() string {
	if p.size == 0 {
		return ""
	}

	sb := strings.Builder{}
	for node := p.head; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%v -> ", node.data))
	}

	return sb.String()
}

func (p *Single[T]) remove(element T) *snode[T] {
	var (
		dummy = snode[T]{}
		prev  = &dummy
		curr  = p.head
	)

	prev.next = curr

	for curr != nil {
		if any(curr.data) == any(element) {
			prev.next = curr.next
			if curr == p.tail {
				p.tail = prev
			}
			p.size--
		} else {
			prev = curr
		}

		curr = curr.next
	}

	return dummy.next
}

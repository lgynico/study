package list

import (
	"github.com/lgynico/algo-go/structure/stack"
)

func IsPalindromicList(head *Node) bool {
	if head == nil {
		return true
	}

	var (
		slow = &Node{Next: head}
		fast = slow
	)

	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}

	var (
		prev *Node
		curr = slow
	)

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	var (
		yes = true
		pl  = head
		pr  = prev
	)

	for pl != nil && pr != nil {
		if pl.Data != pr.Data {
			yes = false
			break
		}
		pl = pl.Next
		pr = pr.Next
	}

	curr = prev
	prev = nil
	for curr != slow {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	slow.Next = prev

	return yes
}

func isPalindromicList(head *Node) bool {
	if head == nil {
		return true
	}

	stack := stack.Array[*Node]{}
	for node := head; node != nil; node = node.Next {
		stack.Push(node)
	}

	for head != nil {
		e, _ := stack.Pop()
		if e.Data != head.Data {
			return false
		}
		head = head.Next
	}

	return true
}

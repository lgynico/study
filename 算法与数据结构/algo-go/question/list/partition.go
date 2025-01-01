package list

import "github.com/lgynico/algo-go/structure/queue"

func Partition(head *Node, element int) *Node {
	if head == nil {
		return nil
	}

	var nodes [6]*Node
	var curr = head
	for curr != nil {
		if curr.Data < element {
			if nodes[0] == nil {
				nodes[0] = curr
				nodes[1] = curr
			} else {
				nodes[1].Next = curr
				nodes[1] = curr
			}
		} else if curr.Data > element {
			if nodes[4] == nil {
				nodes[4] = curr
				nodes[5] = curr
			} else {
				nodes[5].Next = curr
				nodes[5] = curr
			}
		} else {
			if nodes[2] == nil {
				nodes[2] = curr
				nodes[3] = curr
			} else {
				nodes[3].Next = curr
				nodes[3] = curr
			}
		}
		curr = curr.Next
	}

	if nodes[1] != nil {
		nodes[1].Next = nil
	}
	if nodes[3] != nil {
		nodes[3].Next = nil
	}
	if nodes[5] != nil {
		nodes[5].Next = nil
	}

	var (
		newHead *Node
		newTail *Node
	)
	for i := 0; i < len(nodes); i += 2 {
		if newHead == nil {
			newHead = nodes[i]
			newTail = nodes[i+1]
		} else {
			newTail.Next = nodes[i]
			if nodes[i+1] != nil {
				newTail = nodes[i+1]
			}
		}
	}

	return newHead
}

func partition(head *Node, element int) *Node {
	if head == nil {
		return nil
	}

	var (
		lessQueue  = queue.Array[*Node]{}
		equalQueue = queue.Array[*Node]{}
		greatQueue = queue.Array[*Node]{}
	)

	for head != nil {
		if head.Data < element {
			lessQueue.Enqueue(head)
		} else if head.Data > element {
			greatQueue.Enqueue(head)
		} else {
			equalQueue.Enqueue(head)
		}
		head = head.Next
	}

	var dummy = Node{}
	var curr = &dummy
	for !lessQueue.IsEmpty() {
		node, _ := lessQueue.Dequeue()
		curr.Next = node
		curr = node
	}
	for !equalQueue.IsEmpty() {
		node, _ := equalQueue.Dequeue()
		curr.Next = node
		curr = node
	}
	for !greatQueue.IsEmpty() {
		node, _ := greatQueue.Dequeue()
		curr.Next = node
		curr = node
	}

	if curr != nil {
		curr.Next = nil
	}

	return dummy.Next
}

package list

func CenterBehind(head *Node) *Node {
	if head == nil {
		return nil
	}

	var (
		slow = &Node{Next: head}
		fast = slow
	)

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}

	return slow
}

func CenterBefore(head *Node) *Node {
	if head == nil {
		return nil
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

	return slow
}

func center(head *Node, before bool) *Node {
	if head == nil {
		return nil
	}

	size := 0
	for node := head; node != nil; node = node.Next {
		size++
	}

	var cursor int
	if size%2 == 1 {
		cursor = size / 2
	} else {
		cursor = size / 2
		if before {
			cursor--
		}
	}

	node := head
	for cursor > 0 {
		node = node.Next
		cursor--
	}

	return node
}

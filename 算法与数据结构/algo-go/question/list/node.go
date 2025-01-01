package list

import (
	"fmt"
	"strings"

	"github.com/lgynico/algo-go/utils/rands"
)

type Node struct {
	Data int
	Next *Node
}

func RandomList(maxSize, min, max int) *Node {
	var (
		size  = rands.Random(0, maxSize)
		dummy = Node{}
		curr  = &dummy
	)

	for i := 0; i < size; i++ {
		node := &Node{
			Data: rands.Random(min, max),
			Next: nil,
		}
		curr.Next = node
		curr = node
	}

	return dummy.Next
}

func MakeList(elements ...int) *Node {
	var dummy = Node{}
	curr := &dummy

	for _, element := range elements {
		node := &Node{
			Data: element,
			Next: nil,
		}
		curr.Next = node
		curr = node
	}

	return dummy.Next
}

func MakePalindromicList(maxSize, min, max int) *Node {
	var (
		size = rands.Random(0, maxSize)
		arr  = make([]int, size)
		i    = 0
		j    = size - 1
	)

	for i <= j {
		num := rands.Random(min, max)
		arr[i] = num
		arr[j] = num
		i++
		j--
	}

	return MakeList(arr...)
}

func Copy(list *Node) *Node {
	if list == nil {
		return nil
	}

	var dummy = Node{}
	curr := &dummy

	for node := list; node != nil; node = node.Next {
		newNode := &Node{
			Data: node.Data,
			Next: nil,
		}
		curr.Next = newNode
		curr = newNode
	}

	return dummy.Next
}

func ToString(head *Node) string {
	if head == nil {
		return ""
	}

	sb := strings.Builder{}
	for node := head; node != nil; node = node.Next {
		sb.WriteString(fmt.Sprintf("%v -> ", node.Data))
	}

	return sb.String()
}

func Equals(list1, list2 *Node) bool {
	for list1 != nil && list2 != nil {
		if list1.Data != list2.Data {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	return list1 == nil && list2 == nil
}

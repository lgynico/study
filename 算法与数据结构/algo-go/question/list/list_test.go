package list

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/lgynico/algo-go/utils/rands"
)

func TestCenter(t *testing.T) {
	var (
		testCount = 100000
		size      = 1000
		max       = 100
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			list  = RandomList(size, 0, max)
			node1 = CenterBefore(list)
			node2 = center(list, true)
			nodeA = CenterBehind(list)
			nodeB = center(list, false)
		)
		if node1 != node2 || nodeA != nodeB {
			fmt.Println("test erorr.")
			fmt.Println(ToString(list))
			fmt.Println(node1, node2)
			fmt.Println(nodeA, nodeB)
			return
		}
	}
	fmt.Println("test end.")
}

func TestIsPalindromicList(t *testing.T) {
	var (
		testCount = 100000
		size      = 1000
		max       = 100
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			rand = rand.Float64()
			list *Node
		)

		if rand < 0.5 {
			list = RandomList(size, 0, max)
		} else {
			list = MakePalindromicList(size, 0, max)
		}

		result1 := IsPalindromicList(list)
		result2 := isPalindromicList(list)

		if result1 != result2 {
			{
				fmt.Println("test erorr.")
				fmt.Println(result1, result2, ToString(list))
				return
			}
		}
	}
	fmt.Println("test end.")

}

func TestPartition(t *testing.T) {
	var (
		testCount = 100000
		size      = 1000
		max       = 100
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			list    = RandomList(size, 0, max)
			list2   = Copy(list)
			pivot   = rands.Random(0, max)
			result  = Partition(list, pivot)
			result2 = partition(list2, pivot)
		)

		if !Equals(result, result2) {
			fmt.Println("test erorr.")

		}
	}
	fmt.Println("test end.")
}

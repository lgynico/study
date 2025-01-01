package heap

import (
	"fmt"
	"sort"
	"testing"

	"github.com/lgynico/algo-go/utils/arrays"
	"github.com/lgynico/algo-go/utils/rands"
)

func TestOverlapLines(t *testing.T) {
	var (
		testCount = 1_000_000
		size      = 100
		max       = 1000
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			lines   = generateLines(0, max, size)
			result2 = overlapLines(lines)
			result1 = overlapLines(lines)
		)

		if result1 != result2 {
			fmt.Println("test error.")
			fmt.Println(lines)
			fmt.Println(result1, result2)
			return
		}
	}
	fmt.Println("test end.")
}

func TestTopKWinners(t *testing.T) {
	var (
		testCount = 10_000
		size      = 10
		max       = 100
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			users   = arrays.RandomArray(0, max, size)
			oprs    = genBoolArray(len(users))
			k       = rands.Random(1, 10)
			result1 = TopKWinners(users, oprs, k)
			result2 = topKWinners(users, oprs, k)
		)

		if len(result1) != len(result2) {
			fmt.Println("test error.")
			fmt.Println(users)
			fmt.Println(oprs)
			fmt.Println(k)
			fmt.Println(result1)
			fmt.Println(result2)
			return
		}
		for i := 0; i < len(result1); i++ {
			r1 := result1[i]
			r2 := result2[i]
			sort.Ints(r1)
			sort.Ints(r2)
			if !arrays.Equals(r1, r2) {
				fmt.Println("test error.")
				fmt.Println(users)
				fmt.Println(oprs)
				fmt.Println(k)
				fmt.Println(result1)
				fmt.Println(result2)
				return
			}
		}
	}
	fmt.Println("test end.")
}

func TestXxx(t *testing.T) {
	fmt.Println(topKWinners([]int{1, 9, 8, 9, 1}, []bool{true, true, true, false, false}, 2))
}

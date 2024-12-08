package mergesort

import (
	"fmt"
	"testing"

	"github.com/lgynico/algo-go/utils/arrays"
)

func TestSmallSum(t *testing.T) {
	var (
		testCount = 10000
		size      = 100
		max       = 1000
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			arr1    = arrays.RandomArray(0, max, size)
			arr2    = arrays.CopyArray(arr1)
			result1 = SmallSum(arr1)
			result2 = smallSum(arr2)
		)

		if result1 != result2 {
			fmt.Println("test error.")
			fmt.Println(arr1, result1, arr2, result2)
			return
		}
	}
	fmt.Println("test end.")
}

func TestInversePairs(t *testing.T) {
	var (
		testCount = 10000
		size      = 100
		max       = 1000
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			arr1    = arrays.RandomArray(0, max, size)
			arr2    = arrays.CopyArray(arr1)
			result1 = InversePairs(arr1)
			result2 = inversePairs(arr2)
		)

		if result1 != result2 {
			fmt.Println("test error.")
			fmt.Println(arr1, result1, arr2, result2)
			return
		}
	}
	fmt.Println("test end.")
}

func TestTwiceSmall(t *testing.T) {
	var (
		testCount = 10000
		size      = 100
		max       = 1000
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			arr1    = arrays.RandomArray(0, max, size)
			arr2    = arrays.CopyArray(arr1)
			result1 = TwiceSmall(arr1)
			result2 = twiceSmall(arr2)
		)

		if result1 != result2 {
			fmt.Println("test error.")
			fmt.Println(arr1, result1, arr2, result2)
			return
		}
	}
	fmt.Println("test end.")
}

func TestX(t *testing.T) {
	fmt.Println(TwiceSmall([]int{1, 2, 6, 0}))
}

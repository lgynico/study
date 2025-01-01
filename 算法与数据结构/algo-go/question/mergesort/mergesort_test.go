package mergesort

import (
	"fmt"
	"testing"

	"github.com/lgynico/algo-go/utils/arrays"
	"github.com/lgynico/algo-go/utils/rands"
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

func TestCountOfRangeSum(t *testing.T) {
	var (
		testCount = 10000
		size      = 100
		max       = 1000
	)

	fmt.Println("test start.")
	for i := 0; i < testCount; i++ {
		var (
			arr     = arrays.RandomArray(0, max, size)
			lower   = rands.Random(0, max)
			upper   = rands.Random(lower, max)
			result1 = CountOfRangeSum(arrays.CopyArray(arr), lower, upper)
			result2 = countOfRangeSum(arrays.CopyArray(arr), lower, upper)
		)

		if result1 != result2 {
			fmt.Println("test error.")
			fmt.Println(arr)
			fmt.Println(lower, upper)
			fmt.Println(result1, result2)
			return
		}
	}
	fmt.Println("test end.")
}

func TestX(t *testing.T) {
	fmt.Println(CountOfRangeSum([]int{3, 6, 4, 4}, 0, 10))
}

package xor

import (
	"fmt"
	"testing"

	"github.com/lgynico/algo-go/utils/arrays"
	"github.com/lgynico/algo-go/utils/rands"
)

func TestFindOddTimes(t *testing.T) {
	fmt.Println("test begin")
	for i := 0; i < 10000; i++ {
		var (
			arr1    = randomArrayForFindOddTimes(1, 100)
			arr2    = arrays.CopyArray(arr1)
			result1 = FindOddTimes(arr1)
			result2 = findOddTimes(arr2)
		)

		if result1 != result2 {
			fmt.Println(arr1)
			fmt.Println(result1, result2)
			fmt.Println("test error")
			return
		}
	}
	fmt.Println("test success")
}

func TestFindTwoOddTimes(t *testing.T) {
	fmt.Println("test begin")
	for i := 0; i < 10000; i++ {
		var (
			arr1       = randomArrayForFindTwoOddTimes(1, 100)
			arr2       = arrays.CopyArray(arr1)
			ans0, ans1 = findTwoOddTimes(arr1)
			ans2, ans3 = findTwoOddTimes(arr2)
		)

		if !(ans0 == ans2 && ans1 == ans3 || ans0 == ans3 && ans1 == ans2) {
			fmt.Println("test error")
			return
		}
	}
	fmt.Println("test success")
}

func randomArrayForFindOddTimes(min, max int) []int {
	arr := make([]int, 0, 256)
	evenSize := rands.Random(2, 10)
	for i := 0; i < evenSize; i++ {
		num := rands.Random(min, max)
		even := rands.RandomEven(1, 20)
		for j := 0; j < even; j++ {
			arr = append(arr, num)
		}
	}

	num := rands.Random(min, max)
	odd := rands.RandomOdd(1, 20)
	for i := 0; i < odd; i++ {
		arr = append(arr, num)
	}

	arrays.Shuffle(arr)
	return arr
}

func randomArrayForFindTwoOddTimes(min, max int) []int {
	arr := make([]int, 0, 256)
	evenSize := rands.Random(2, 10)
	for i := 0; i < evenSize; i++ {
		num := rands.Random(min, max)
		even := rands.RandomEven(1, 20)
		for j := 0; j < even; j++ {
			arr = append(arr, num)
		}
	}

	for i := 0; i < 2; i++ {
		num := rands.Random(min, max)
		odd := rands.RandomOdd(1, 20)
		for i := 0; i < odd; i++ {
			arr = append(arr, num)
		}
	}

	arrays.Shuffle(arr)
	return arr
}

package arrays

import (
	"math/rand"

	"github.com/lgynico/algo-go/utils/constraints"
	"github.com/lgynico/algo-go/utils/rands"
)

func RandomArray(min, max, size int) []int {
	if size <= 0 {
		return []int{}
	}

	arr := make([]int, 0, size)
	for i := 0; i < size; i++ {
		num := min + rand.Intn(max-min)
		arr = append(arr, num)
	}

	return arr
}

func CopyArray[T any](arr []T) []T {
	temp := make([]T, len(arr))
	copy(temp, arr)
	return temp
}

func Equals[T constraints.Comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func Shuffle[T any](arr []T) {
	for i := 0; i < len(arr); i++ {
		j := rands.Random(0, len(arr)-1)
		if i == j {
			return
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}

func MaxMin[T constraints.Comparable](arr []T) (max, min T) {
	if len(arr) == 0 {
		return
	}

	max, min = arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return
}

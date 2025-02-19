package sort

import "github.com/lgynico/algo-go/utils/constraints"

func Bubble[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

package sort

import "github.com/lgynico/algo-go/utils/constraints"

func Insertion[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] >= arr[j-1] {
				break
			}

			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

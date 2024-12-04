package sort

import "github.com/lgynico/algo-go/utils/constraints"

func Selection[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

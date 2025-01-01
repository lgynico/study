package sort

import "github.com/lgynico/algo-go/utils/constraints"

func Heap[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr)-1)
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i-1)
	}
}

func heapify[T constraints.Comparable](arr []T, i, j int) {
	for i <= j {
		child := 2*i + 1
		if child > j {
			break
		}
		if child+1 <= j && arr[child+1] > arr[child] {
			child++
		}
		if arr[child] <= arr[i] {
			break
		}

		arr[i], arr[child] = arr[child], arr[i]
		i = child
	}
}

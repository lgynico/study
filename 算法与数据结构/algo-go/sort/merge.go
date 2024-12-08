package sort

import "github.com/lgynico/algo-go/utils/constraints"

func Merge[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	mergeSort(arr, 0, len(arr)-1)
}

func MergeAdvance[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	step := 1
	for step < len(arr) {
		left := 0
		for left < len(arr) {
			var (
				mid   = left + step - 1
				right = mid + step
			)
			if mid >= len(arr) {
				mid = len(arr) - 1
			}
			if right >= len(arr) {
				right = len(arr) - 1
			}

			merge(arr, left, mid, right)
			left = right + 1
		}
		step <<= 1
	}
}

func mergeSort[T constraints.Comparable](arr []T, left, right int) {
	if left < right {
		mid := left + (right-left)>>1
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge[T constraints.Comparable](arr []T, left, mid, right int) {
	var (
		temp                = make([]T, right-left+1)
		currLeft, maxLeft   = left, mid
		currRight, maxRight = mid + 1, right
		i                   = 0
	)

	for currLeft <= maxLeft && currRight <= maxRight {
		if arr[currLeft] <= arr[currRight] {
			temp[i] = arr[currLeft]
			currLeft++
		} else {
			temp[i] = arr[currRight]
			currRight++
		}
		i++
	}

	for currLeft <= maxLeft {
		temp[i] = arr[currLeft]
		currLeft++
		i++
	}

	for currRight <= maxRight {
		temp[i] = arr[currRight]
		currRight++
		i++
	}

	copy(arr[left:right+1], temp)
}

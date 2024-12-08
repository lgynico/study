package sort

import (
	"github.com/lgynico/algo-go/structure/stack"
	"github.com/lgynico/algo-go/utils/constraints"
	"github.com/lgynico/algo-go/utils/rands"
)

func Quick[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	quickSort(arr, 0, len(arr)-1)
}

func QuickAdvanced[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	var (
		pl, pr = partition(arr, 0, len(arr)-1)
		stack  = stack.Array[[]int]{}
	)
	stack.Push([]int{0, pl})
	stack.Push([]int{pr, len(arr) - 1})

	for !stack.IsEmpty() {
		var (
			p, _  = stack.Pop()
			left  = p[0]
			right = p[1]
		)

		pl, pr = partition(arr, left, right)
		if pl > left {
			stack.Push([]int{left, pl})
		}
		if pr < right {
			stack.Push([]int{pr, right})
		}
	}
}

func quickSort[T constraints.Comparable](arr []T, left, right int) {
	if left >= right {
		return
	}

	pl, pr := partition(arr, left, right)
	quickSort(arr, left, pl)
	quickSort(arr, pr, right)
}

func partition[T constraints.Comparable](arr []T, left, right int) (int, int) {
	ri := rands.Random(left, right)
	arr[right], arr[ri] = arr[ri], arr[right]

	var (
		x  = arr[right]
		pl = left - 1
		pr = right + 1
		i  = left
	)

	for i < pr {
		if arr[i] < x {
			arr[i], arr[pl+1] = arr[pl+1], arr[i]
			pl++
			i++
		} else if arr[i] > x {
			arr[i], arr[pr-1] = arr[pr-1], arr[i]
			pr--
		} else {
			i++
		}
	}

	return pl, pr
}

func QuickV1[T constraints.Comparable](arr []T) {
	if len(arr) <= 1 {
		return
	}

	quickSortV1(arr, 0, len(arr)-1)
}

func quickSortV1[T constraints.Comparable](arr []T, left, right int) {
	if left >= right {
		return
	}

	pl, pr := partitionV1(arr, left, right)
	quickSortV1(arr, left, pl)
	quickSortV1(arr, pr, right)
}

func partitionV1[T constraints.Comparable](arr []T, left, right int) (int, int) {
	var (
		x  = arr[right]
		pl = left - 1
		pr = right + 1
		i  = left
	)

	for i < pr {
		if arr[i] < x {
			arr[i], arr[pl+1] = arr[pl+1], arr[i]
			pl++
			i++
		} else if arr[i] > x {
			arr[i], arr[pr-1] = arr[pr-1], arr[i]
			pr--
		} else {
			i++
		}
	}

	return pl, pr
}

func netherlandFlag1[T constraints.Comparable](arr []T, x T) int {
	if len(arr) == 0 {
		return -1
	}

	p := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] <= x {
			arr[i], arr[p+1] = arr[p+1], arr[i]
			p++
		}
	}
	return p
}

func netherlandFlag2[T constraints.Comparable](arr []T, x T) (int, int) {
	if len(arr) == 0 {
		return -1, -1
	}

	var (
		pl = -1
		pr = len(arr)
		i  = 0
	)

	for i < pr {
		if arr[i] < x {
			arr[i], arr[pl+1] = arr[pl+1], arr[i]
			pl++
			i++
		} else if arr[i] > x {
			arr[i], arr[pr-1] = arr[pr-1], arr[i]
			pr--
		} else {
			i++
		}
	}

	return pl, pr
}

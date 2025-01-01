package sort

import (
	"github.com/lgynico/algo-go/utils/arrays"
)

func Count(arr []int) {
	if len(arr) <= 1 {
		return
	}

	max, min := arrays.MaxMin[int](arr)
	for _, i := range arr {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	if min < 0 {
		max -= min
		for i := 0; i < len(arr); i++ {
			arr[i] -= min
		}
	}

	helper := make([]int, max+1)
	for _, i := range arr {
		helper[i]++
	}

	var i int
	for j, num := range helper {
		for num > 0 {
			arr[i] = j
			i++
			num--
		}
	}

	if min < 0 {
		for i := 0; i < len(arr); i++ {
			arr[i] += min
		}
	}
}

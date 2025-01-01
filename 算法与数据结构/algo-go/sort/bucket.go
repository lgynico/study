package sort

import (
	"math"

	"github.com/lgynico/algo-go/structure/queue"
	"github.com/lgynico/algo-go/utils/arrays"
)

func Bucket(arr []int) {
	if len(arr) <= 1 {
		return
	}

	var (
		digit    = 1
		max, min = arrays.MaxMin[int](arr)
	)

	if min < 0 {
		max -= min
		for i := 0; i < len(arr); i++ {
			arr[i] -= min
		}
	}

	for max >= 10 {
		digit++
		max /= 10
	}

	buckets := make([]queue.Array[int], 10)
	for i := 0; i < digit; i++ {
		for _, num := range arr {
			index := int(float64(num)/math.Pow10(i)) % 10
			buckets[index].Enqueue(num)
		}
		var j int
		for k := 0; k < len(buckets); k++ {
			for !buckets[k].IsEmpty() {
				arr[j], _ = buckets[k].Dequeue()
				j++
			}
		}
	}

	if min < 0 {
		for i := 0; i < len(arr); i++ {
			arr[i] += min
		}
	}
}

func BucketAdvance(arr []int) {
	if len(arr) <= 1 {
		return
	}

	var (
		max, min = arrays.MaxMin(arr)
		digit    = 1
	)

	if min < 0 {
		max -= min
		for i := 0; i < len(arr); i++ {
			arr[i] -= min
		}
	}

	for max >= 10 {
		digit++
		max /= 10
	}

	var (
		count  = make([]int, 10)
		helper = make([]int, len(arr))
	)

	for i := 0; i < digit; i++ {
		for j := 0; j < len(count); j++ {
			count[j] = 0
		}

		for _, num := range arr {
			index := int(float64(num)/math.Pow10(i)) % 10
			count[index]++
		}

		for j := 1; j < len(count); j++ {
			count[j] += count[j-1]
		}

		for j := len(arr) - 1; j >= 0; j-- {
			index := int(float64(arr[j])/math.Pow10(i)) % 10
			helper[count[index]-1] = arr[j]
			count[index]--
		}

		copy(arr, helper)
	}

	if min < 0 {
		for i := 0; i < len(arr); i++ {
			arr[i] += min
		}
	}
}

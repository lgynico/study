package mergesort

func InversePairs(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return inversePairsSort(arr, 0, len(arr)-1)
}

func inversePairsSort(arr []int, left, right int) int {
	var result int
	if left < right {
		mid := left + ((right - left) >> 1)
		result += inversePairsSort(arr, left, mid)
		result += inversePairsSort(arr, mid+1, right)
		result += inversePairsMerge(arr, left, mid, right)
	}
	return result
}

func inversePairsMerge(arr []int, left, mid, right int) int {
	var (
		temp      = make([]int, right-left+1)
		result    int
		i         int
		currLeft  = left
		maxLeft   = mid
		currRight = mid + 1
		maxRight  = right
	)

	for currLeft <= maxLeft && currRight <= maxRight {
		if arr[currLeft] <= arr[currRight] {
			temp[i] = arr[currLeft]
			currLeft++
		} else {
			result += maxLeft - currLeft + 1
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

	return result
}

func inversePairs(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	result := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				result++
			}
		}
	}
	return result
}

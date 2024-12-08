package mergesort

func SmallSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return smallSumSort(arr, 0, len(arr)-1)
}

func smallSumSort(arr []int, left, right int) int {
	result := 0
	if left < right {
		mid := left + (right-left)>>1
		result += smallSumSort(arr, left, mid)
		result += smallSumSort(arr, mid+1, right)
		result += smallSumMerge(arr, left, mid, right)
	}

	return result
}

func smallSumMerge(arr []int, left, mid, right int) int {
	var (
		result    int
		temp      = make([]int, right-left+1)
		currLeft  = left
		maxLeft   = mid
		currRight = mid + 1
		maxRight  = right
		i         int
	)

	for currLeft <= maxLeft && currRight <= maxRight {
		if arr[currLeft] < arr[currRight] {
			temp[i] = arr[currLeft]
			result += (maxRight - currRight + 1) * arr[currLeft]
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

	return result
}

func smallSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	var result int
	for i := len(arr) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[i] {
				result += arr[j]
			}
		}
	}
	return result
}

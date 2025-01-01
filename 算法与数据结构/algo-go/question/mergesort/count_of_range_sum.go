package mergesort

func CountOfRangeSum(arr []int, lower, upper int) int {
	if len(arr) == 0 {
		return 0
	}

	preSum := make([]int, len(arr))
	preSum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		preSum[i] = arr[i] + preSum[i-1]
	}

	return countOfRangeSumSort(preSum, 0, len(preSum)-1, lower, upper)
}

func countOfRangeSumSort(arr []int, left, right, lower, upper int) int {
	if left == right {
		if arr[left] >= lower && arr[left] <= upper {
			return 1
		}
		return 0
	}

	var result int
	mid := left + ((right - left) >> 1)

	result += countOfRangeSumSort(arr, left, mid, lower, upper)
	result += countOfRangeSumSort(arr, mid+1, right, lower, upper)
	result += countOfRangeSumMerge(arr, left, mid, right, lower, upper)

	return result
}

func countOfRangeSumMerge(arr []int, left, mid, right int, lower, upper int) int {
	var (
		result   int
		winLeft  = left
		winRight = left
	)

	for i := mid + 1; i <= right; i++ {
		min := arr[i] - upper
		max := arr[i] - lower
		for winLeft <= mid && arr[winLeft] < min {
			winLeft++
		}
		for winRight <= mid && arr[winRight] <= max {
			winRight++
		}
		result += winRight - winLeft
	}

	var (
		temp      = make([]int, right-left+1)
		currLeft  = left
		maxLeft   = mid
		currRight = mid + 1
		maxRight  = right
		i         int
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

	return result
}

func countOfRangeSum(arr []int, lower, upper int) int {
	var result int

	for i := 0; i < len(arr); i++ {
		for j := 0; j <= i; j++ {
			var sum int
			for k := j; k <= i; k++ {
				sum += arr[k]
			}
			if sum >= lower && sum <= upper {
				result++
			}
		}
	}

	return result
}

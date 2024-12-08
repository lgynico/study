package mergesort

func TwiceSmall(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return twiceSmallSort(arr, 0, len(arr)-1)
}

func twiceSmallSort(arr []int, left, right int) int {
	var result int
	if left < right {
		mid := left + (right-left)>>1
		result += twiceSmallSort(arr, left, mid)
		result += twiceSmallSort(arr, mid+1, right)
		result += twiceSmallMerge(arr, left, mid, right)
	}
	return result
}

func twiceSmallMerge(arr []int, left, mid, right int) int {
	var (
		result int
		i      = mid
		j      = right
	)

	for i >= left && j >= mid+1 {
		if arr[i] > arr[j]*2 {
			result += j - mid
			i--
		} else {
			j--
		}
	}

	var (
		temp      = make([]int, right-left+1)
		index     int
		currLeft  = left
		maxLeft   = mid
		currRight = mid + 1
		maxRight  = right
	)

	for currLeft <= maxLeft && currRight <= maxRight {
		if arr[currLeft] <= arr[currRight] {
			temp[index] = arr[currLeft]
			currLeft++
		} else {
			temp[index] = arr[currRight]
			currRight++
		}
		index++
	}

	for currLeft <= maxLeft {
		temp[index] = arr[currLeft]
		currLeft++
		index++
	}

	for currRight <= maxRight {
		temp[index] = arr[currRight]
		currRight++
		index++
	}

	copy(arr[left:right+1], temp)

	return result
}

func twiceSmall(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	var result int
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j]*2 {
				result++
			}
		}
	}
	return result
}

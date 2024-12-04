package xor

func FindOddTimes(arr []int) int {
	xor := 0
	for _, num := range arr {
		xor ^= num
	}
	return xor
}

func findOddTimes(arr []int) int {
	m := map[int]int{}
	for _, num := range arr {
		m[num]++
	}

	for num, count := range m {
		if (count & 1) == 1 {
			return num
		}
	}

	return -1
}

package xor

func FindTwoOddTimes(arr []int) (int, int) {
	xor := 0
	for _, num := range arr {
		xor ^= num
	}

	oneBit := xor ^ -xor
	ans := 0
	for _, num := range arr {
		if num&oneBit == oneBit {
			ans ^= num
		}
	}

	return ans, xor ^ ans
}

func findTwoOddTimes(arr []int) (int, int) {
	var (
		m   = map[int]int{}
		ans []int
	)

	for _, num := range arr {
		m[num]++
	}

	for k, v := range arr {
		if v%2 != 0 {
			ans = append(ans, k)
		}
	}

	if len(ans) > 2 {
		return ans[0], ans[1]
	}
	return 0, 0
}

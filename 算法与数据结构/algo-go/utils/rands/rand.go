package rands

import "math/rand"

func Random(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomOdd(min, max int) int {
	num := Random(min, max)
	if num%2 == 1 {
		return num
	}

	if num == min {
		return num + 1
	}
	if num == max {
		return num - 1
	}

	if rand.Float64() < .5 {
		return num + 1
	}

	return num - 1
}

func RandomEven(min, max int) int {
	num := Random(min, max)
	if num%2 == 0 {
		return num
	}

	if num == min {
		return num + 1
	}
	if num == max {
		return num - 1
	}

	if rand.Float64() < .5 {
		return num + 1
	}

	return num - 1
}

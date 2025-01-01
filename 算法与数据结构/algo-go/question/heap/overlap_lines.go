package heap

import (
	"math"
	"math/rand"
	"sort"

	"github.com/lgynico/algo-go/structure/heap"
	"github.com/lgynico/algo-go/utils/rands"
)

func OverlapLines(lines [][2]int) int {
	if len(lines) == 0 {
		return 0
	}

	sort.Slice(lines, func(i, j int) bool { return lines[i][0] <= lines[j][0] })

	var (
		h      = heap.NewArray(func(a, b int) int { return a - b })
		result int
	)

	for _, line := range lines {
		for !h.IsEmpty() && h.Peek() <= line[0] {
			h.Pop()
		}

		h.Push(line[1])

		if h.Size() > result {
			result = h.Size()
		}
	}

	return result
}

func overlapLines(lines [][2]int) int {
	if len(lines) == 0 {
		return 0
	}

	var (
		min    = math.MaxInt
		max    = math.MinInt
		result int
	)

	for _, line := range lines {
		if line[0] < min {
			min = line[0]
		}
		if line[0] > max {
			max = line[0]
		}
	}

	for i := float64(min) + 0.5; i < float64(max); i++ {
		count := 0
		for _, line := range lines {
			if float64(line[0]) < i && float64(line[1]) > i {
				count++
			}
		}

		if count > result {
			result = count
		}
	}

	return result
}

func generateLines(min, max, size int) [][2]int {
	lines := make([][2]int, 0, rands.Random(0, size))
	for i := 0; i < len(lines); i++ {
		start := rands.Random(min, max)
		end := rands.Random(min, max)
		if start > end {
			start, end = end, start
		}
		if start == end {
			if start > 0 {
				start--
			} else if end < max {
				end++
			} else if rand.Float64() < .5 {
				start--
			} else {
				end++
			}
		}
		lines = append(lines, [2]int{start, end})
	}

	return lines
}

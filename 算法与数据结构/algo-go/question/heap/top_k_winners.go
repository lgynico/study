package heap

import (
	"math/rand"
	"sort"

	"github.com/lgynico/algo-go/structure/heap"
)

func TopKWinners(users []int, isBuy []bool, k int) [][]int {
	helper := newTopKHelper(k)
	for i := 0; i < len(users); i++ {
		helper.Process(i, users[i], isBuy[i])
	}
	return helper.Results()
}

func topKWinners(users []int, isBuy []bool, k int) [][]int {
	var (
		results    = make([][]int, 0, len(users))
		customers  = make(map[int]*customer, k)
		winners    = make([]*customer, 0, k)
		candidates = make([]*customer, 0, k)
	)

	for i := 0; i < len(users); i++ {
		var (
			user  = users[i]
			buy   = isBuy[i]
			c, ok = customers[user]
		)

		if !ok {
			c = &customer{id: user, num: 0, time: 0}
			if buy {
				customers[user] = c
			}
		}

		if buy {
			c.num++
			if _, ok := inArray(c, winners); !ok {
				if _, ok = inArray(c, candidates); !ok {
					c.time = i
					if len(winners) < k {
						winners = append(winners, c)
					} else {
						candidates = append(candidates, c)
					}
				}
			}
		} else {
			c.num--
			if c.num == 0 {
				delete(customers, c.id)
				if i, ok := inArray(c, winners); ok {
					winners = append(winners[:i], winners[i+1:]...)
				} else {
					if i, ok := inArray(c, candidates); ok {
						candidates = append(candidates[:i], candidates[i+1:]...)
					}
				}
			}
		}

		if len(candidates) != 0 {
			sort.Slice(winners, func(i, j int) bool {
				wi := winners[i]
				wj := winners[j]
				if wi.num != wj.num {
					return wi.num < wj.num
				}
				return wi.time < wj.time
			})

			sort.Slice(candidates, func(i, j int) bool {
				ci := candidates[i]
				cj := candidates[j]
				if ci.num != cj.num {
					return ci.num > cj.num
				}
				return ci.time < cj.time
			})

			topCandidate := candidates[0]
			if len(winners) < k {
				topCandidate.time = i
				winners = append(winners, topCandidate)
				candidates = candidates[1:]
			} else {
				botWinner := winners[0]

				if topCandidate.num > botWinner.num {
					topCandidate.time = i
					botWinner.time = i

					candidates = candidates[1:]
					winners = winners[1:]

					candidates = append(candidates, botWinner)
					winners = append(winners, topCandidate)
				}
			}
		}

		result := make([]int, 0, len(winners))
		for _, c := range winners {
			result = append(result, c.id)
		}
		results = append(results, result)
	}

	return results
}

func inArray[T any](element T, arr []T) (int, bool) {
	for i, e := range arr {
		if any(e) == any(element) {
			return i, true
		}
	}
	return -1, false
}

type (
	customer struct {
		id   int
		num  int
		time int
	}
	topKHelper struct {
		k          int
		customers  map[int]*customer
		winners    heap.Advance[*customer]
		candidates heap.Advance[*customer]
		results    [][]int
	}
)

func newTopKHelper(k int) topKHelper {
	return topKHelper{
		k:         k,
		customers: make(map[int]*customer),
		winners: heap.NewAdvance(func(a, b *customer) int {
			if a.num != b.num {
				return a.num - b.num
			}
			return a.time - b.time
		}),
		candidates: heap.NewAdvance(func(a, b *customer) int {
			if a.num != b.num {
				return b.num - a.num
			}
			return a.time - b.time
		}),
	}
}

func (p *topKHelper) Process(i, user int, isBuy bool) {
	c, ok := p.customers[user]
	if !ok {
		c = &customer{id: user, num: 0, time: 0}
		if isBuy {
			p.customers[user] = c
		}
	}

	if isBuy {
		c.num++
		if p.winners.Contains(c) || p.candidates.Contains(c) {
			if p.winners.Contains(c) {
				p.winners.Resign(c)
			} else {
				p.candidates.Resign(c)
			}
		} else {
			c.time = i
			if p.winners.Size() < p.k {
				p.winners.Push(c)
			} else {
				p.candidates.Push(c)
			}
		}
	} else {
		c.num--
		if c.num == 0 {
			p.winners.Remove(c)
			p.candidates.Remove(c)
			delete(p.customers, c.id)
		} else {
			if p.winners.Contains(c) {
				p.winners.Resign(c)
			} else {
				p.candidates.Resign(c)
			}
		}
	}

	p.check(i)
	p.results = append(p.results, p.getWinners())
}

func (p *topKHelper) Results() [][]int {
	return p.results
}

func (p *topKHelper) check(i int) {
	if p.candidates.IsEmpty() {
		return
	}

	if p.winners.Size() < p.k {
		c := p.candidates.Pop()
		c.time = i
		p.winners.Push(c)
		return
	}

	bottomWinner := p.winners.Peek()
	topCandidate := p.candidates.Peek()
	if bottomWinner.num < topCandidate.num {
		p.winners.Pop()
		p.candidates.Pop()

		bottomWinner.time = i
		topCandidate.time = i

		p.winners.Push(topCandidate)
		p.candidates.Push(bottomWinner)
	}
}

func (p *topKHelper) getWinners() []int {
	result := make([]int, 0, p.k)
	for _, e := range p.winners.Elements() {
		result = append(result, e.id)
	}
	return result
}

func genBoolArray(size int) []bool {
	arr := make([]bool, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Float64() < .75)
	}
	return arr
}

package poker

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/lgynico/algo-go/utils/arrays"
	"github.com/lgynico/algo-go/utils/rands"
)

func TestPoker1(t *testing.T) {
	fmt.Println("test start")
	for i := 0; i < 10000; i++ {
		input := generatePokerHands(false)
		input2 := arrays.CopyArray(input)
		result := Poker01(input)
		result2 := poker01(input2)
		if result != result2 {
			fmt.Println("test error:", PokerHandsToString(input), result, result2)
			return
		}
	}
	fmt.Println("test end")
}

func TestPoker2(t *testing.T) {
	fmt.Println("test start")
	for i := 0; i < 10000; i++ {
		input := generatePokerHands(true)
		input2 := arrays.CopyArray(input)
		result := Poker02(input)
		result2 := poker02(input2)
		if result != result2 {
			fmt.Println("test error:", PokerHandsToString(input), result, result2)
			return
		}
	}
	fmt.Println("test end")
}

func TestTexas(t *testing.T) {
	fmt.Println("test start")
	for i := 0; i < 1; i++ {
		self := generatePokerHands2(2)
		board := generatePokerHands2(5)
		result := Texas(self, board)
		fmt.Println(PokerHandsToString(self), PokerHandsToString(board), PokerHandsString(result))
		// if result != result2 {
		// 	fmt.Println("test error:", PokerToString(input), result, result2)
		// 	return
		// }
	}
	fmt.Println("test end")
}

func TestPoker(t *testing.T) {
	fmt.Println(Poker02([]int{0x10e, 0x10d, 0x30c, Poker_JokerLittle, Poker_JokerBig}))
	fmt.Println(poker02([]int{0x10e, 0x10d, 0x30c, Poker_JokerLittle, Poker_JokerBig}))
	// fmt.Println(poker02([]int{0x108, 0x10e, 0x303, Poker_JokerLittle, Poker_JokerBig}))
}

var (
	pokers = []int{
		0x102, 0x103, 0x104, 0x105, 0x106, 0x107, 0x108, 0x109, 0x10a, 0x10b, 0x10c, 0x10d, 0x10e,
		0x202, 0x203, 0x204, 0x205, 0x206, 0x207, 0x208, 0x209, 0x20a, 0x20b, 0x20c, 0x20d, 0x20e,
		0x302, 0x303, 0x304, 0x305, 0x306, 0x307, 0x308, 0x309, 0x30a, 0x30b, 0x30c, 0x30d, 0x30e,
		0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e,
	}

	jockers = []int{Poker_JokerLittle, Poker_JokerBig}

	all = []int{
		0x102, 0x103, 0x104, 0x105, 0x106, 0x107, 0x108, 0x109, 0x10a, 0x10b, 0x10c, 0x10d, 0x10e,
		0x202, 0x203, 0x204, 0x205, 0x206, 0x207, 0x208, 0x209, 0x20a, 0x20b, 0x20c, 0x20d, 0x20e,
		0x302, 0x303, 0x304, 0x305, 0x306, 0x307, 0x308, 0x309, 0x30a, 0x30b, 0x30c, 0x30d, 0x30e,
		0x402, 0x403, 0x404, 0x405, 0x406, 0x407, 0x408, 0x409, 0x40a, 0x40b, 0x40c, 0x40d, 0x40e,
		Poker_JokerLittle, Poker_JokerBig,
	}
)

func generatePokerHands(withJocker bool) []int {
	var (
		pokerHands = map[int]struct{}{}
		jokerNum   int
	)

	if withJocker {
		jokerNum = 1 + rand.Intn(2)
		for i := 0; i < jokerNum; i++ {
			pokerHands[jockers[i]] = struct{}{}
		}
	}

	for i := 0; i < 5-jokerNum; i++ {
		for {
			index := rands.Random(0, len(pokers)-1)
			if _, ok := pokerHands[pokers[index]]; !ok {
				pokerHands[pokers[index]] = struct{}{}
				break
			}
		}
	}

	var cards []int
	for card := range pokerHands {
		cards = append(cards, card)
	}

	return cards
}

func generatePokerHands2(num int, excludes ...int) []int {
	var (
		pokerHands    = map[int]struct{}{}
		excludePokers = map[int]struct{}{}
	)

	for _, ex := range excludes {
		excludePokers[ex] = struct{}{}
	}

	for i := 0; i < num; i++ {
		for {
			index := rands.Random(0, len(all)-1)
			poker := all[index]
			if _, ok := excludePokers[poker]; !ok {
				if _, ok1 := pokerHands[poker]; !ok1 {
					pokerHands[poker] = struct{}{}
					break
				}
			}
		}
	}

	keys := make([]int, 0, len(pokerHands))
	for k := range pokerHands {
		keys = append(keys, k)
	}
	return keys
}

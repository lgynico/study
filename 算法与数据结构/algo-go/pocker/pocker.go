package poker

import (
	"sort"
	"strconv"
)

func isFlushKind(input []int) (int, bool) {

	isFlush := isFlush(input)
	total := 0
	for i := 0; i < len(input); i++ {
		total += input[i] & Mask_PokerNumber
	}

	var temp []int
	for _, card := range input {
		temp = append(temp, card&Mask_PokerNumber)
	}
	sort.Ints(temp)

	isStraight := true
	for i := 1; i < len(temp); i++ {
		if isStraight {
			diff := temp[i] - temp[i-1]
			if i != len(temp)-1 {
				isStraight = diff == 1
			} else {
				isStraight = diff == 1 || diff == 9
			}

		}
	}

	if isFlush {
		if total == 60 {
			return PokerHands_RoyalFlush, true
		}
		if isStraight {
			return PokerHands_StraightFlush, true
		}
		return PokerHands_Flush, true
	}

	if isStraight {
		return PokerHands_Straight, true
	}

	return PokerHands_HighCard, false
}

func isPairsKind(input []int) (int, bool) {
	cards := [15]int{}
	for _, card := range input {
		cards[card&Mask_PokerNumber]++
	}

	hasThree := false
	hasPair := false
	for i := 1; i < len(cards); i++ {
		if cards[i] == 4 {
			return PokerHands_FourOfTheKind, true
		}
		if cards[i] == 3 {
			if hasPair {
				return PokerHands_FullHouse, true
			}
			hasThree = true
		}
		if cards[i] == 2 {
			if hasThree {
				return PokerHands_FullHouse, true
			}
			if hasPair {
				return PokerHands_TwoPairs, true
			}
			hasPair = true
		}
	}

	if hasPair {
		return PokerHands_Pair, true
	}
	if hasThree {
		return PokerHands_ThreeOfAKind, true
	}

	return PokerHands_HighCard, false
}

func isFlush(input []int) bool {
	suit := input[0] & Mask_PokerSuit

	for i := 1; i < len(input); i++ {
		if input[i]&Mask_PokerSuit != suit {
			return false
		}
	}

	return true
}

func isStraight(input []int) bool {
	var (
		min    = 0xd
		max    = 0x2
		hasAce bool
		flags  int
	)

	for _, card := range input {
		number := card & Mask_PokerNumber
		if (flags>>number)&0x1 == 0x1 {
			return false
		}

		flags |= 1 << number

		if number == 0xe {
			hasAce = true
			continue
		}

		if number < min {
			min = number
		}
		if number > max {
			max = number
		}

	}

	if !hasAce {
		return max-min <= 4
	}

	return max <= 5 || min >= 10
}

func PokerHandsToString(input []int) string {
	var result string
	for _, card := range input {
		result += PokerToString(card) + " "
	}

	return "[" + result + "]"
}

func PokerToString(card int) string {
	var result string
	if card == Poker_JokerLittle {
		return "小王"
	}
	if card == Poker_JokerBig {
		return "大王"
	}

	switch card & Mask_PokerSuit {
	case PokerSuit_Diamond:
		result += "方块"
	case PokerSuit_Club:
		result += "梅花"
	case PokerSuit_Heart:
		result += "红心"
	case PokerSuit_Spade:
		result += "黑桃"
	}

	switch card & Mask_PokerNumber {
	case 0xa:
		result += "10"
	case 0xb:
		result += "J"
	case 0xc:
		result += "Q"
	case 0xd:
		result += "K"
	case 0xe:
		result += "A"
	default:
		result += strconv.Itoa(card & Mask_PokerNumber)
	}
	return result
}

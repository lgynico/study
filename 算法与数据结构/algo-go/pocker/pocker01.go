package poker

import "sort"

// 算牌型
// 请完成函数f，输入的5个数字代表5张牌，含义如下：
// 0x102,0x103,0x104,0x105,0x106,0x107,0x108,0x109,0x10a,0x10b,0x10c,0x10d,0x10e分别代表方块2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x202,0x203,0x204,0x205,0x206,0x207,0x208,0x209,0x20a,0x20b,0x20c,0x20d,0x20e分别代表梅花2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x302,0x303,0x304,0x305,0x306,0x307,0x308,0x309,0x30a,0x30b,0x30c,0x30d,0x30e分别代表红桃2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x402,0x403,0x404,0x405,0x406,0x407,0x408,0x409,0x40a,0x40b,0x40c,0x40d,0x40e分别代表黑桃2,3,4,5,6,7,8,9,10,J,Q,K,A
// 返回的数字含义如下：
// 1、皇家同花顺：如果花色一样，数字分别是10,J,Q,K,A
// 2、同花顺：如果花色一样，数字是连续的，皇家同花顺除外，例如[0x109,0x10a,0x10b,0x10c,0x10d],[0x10e,0x102,0x103,0x104,0x105]
// 3、金刚：其中4张牌数字一样
// 4、葫芦：其中3张牌数字一样，另外2张牌数字一样
// 5、同花：花色一样，数字不连续
// 6、顺子：数字是连续，花色不一样
// 7、三条：其中3张牌数字一样，另外2张牌数字不一样
// 8、两对：其中2张牌数字一样，另外其中2张牌数字一样，最后一张数字不一样
// 9、一对：其中2张牌数字一样，另外数字不一样
// 10、高牌：什么都不是

func Poker01(input []int) int {
	if kind, ok := isFlushKind(input); ok {
		return kind
	}

	if kind, ok := isPairsKind(input); ok {
		return kind
	}

	return PokerHands_HighCard
}

func poker01(input []int) int {

	// 皇家同花顺
	isSameFlower := true
	suit := input[0] & Mask_PokerSuit
	for _, card := range input {
		if card&Mask_PokerSuit != suit {
			isSameFlower = false
			break
		}
	}

	var tmp []int
	for _, card := range input {
		tmp = append(tmp, card&Mask_PokerNumber)
	}
	sort.Ints(tmp)

	isTJQKA := true
	for i := 0; i < 5; i++ {
		if tmp[i] != 0xa+i {
			isTJQKA = false
			break
		}
	}

	if isSameFlower && isTJQKA {
		return PokerHands_RoyalFlush
	}

	// 同花顺
	isSunzi := true
	for i := 1; i < len(tmp); i++ {
		if isSunzi {
			diff := tmp[i] - tmp[i-1]
			if i != len(tmp)-1 {
				isSunzi = diff == 1
			} else {
				isSunzi = diff == 1 || diff == 9
			}

		}
	}

	if isSameFlower && isSunzi {
		return PokerHands_StraightFlush
	}

	// 四条
	var temp = [15]int{}
	for _, card := range input {
		temp[card&Mask_PokerNumber]++
	}

	for _, num := range temp {
		if num >= 4 {
			return PokerHands_FourOfTheKind
		}
	}

	// 三带二
	hasThree := false
	hasPair := false
	for _, num := range temp {
		if num == 3 {
			hasThree = true
		}
		if num == 2 {
			hasPair = true
		}
	}

	if hasThree && hasPair {
		return PokerHands_FullHouse
	}

	// 同花
	if isSameFlower {
		return PokerHands_Flush
	}

	// 顺子
	if isSunzi {
		return PokerHands_Straight
	}

	// 三条
	if hasThree {
		return PokerHands_ThreeOfAKind
	}

	// 两对
	hasPair1 := false
	hasPair2 := false
	for _, num := range temp {
		if num == 2 {
			if !hasPair1 {
				hasPair1 = true
			} else {
				hasPair2 = true
			}
		}
	}

	if hasPair1 && hasPair2 {
		return PokerHands_TwoPairs
	}

	// 一对
	if hasPair {
		return PokerHands_Pair
	}

	return PokerHands_HighCard
}

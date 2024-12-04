package poker

import "sort"

// 赖子算牌型
// 请完成函数f，输入的5个数字代表5张牌，含义如下：
// 0x102,0x103,0x104,0x105,0x106,0x107,0x108,0x109,0x10a,0x10b,0x10c,0x10d,0x10e分别代表方块2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x202,0x203,0x204,0x205,0x206,0x207,0x208,0x209,0x20a,0x20b,0x20c,0x20d,0x20e分别代表梅花2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x302,0x303,0x304,0x305,0x306,0x307,0x308,0x309,0x30a,0x30b,0x30c,0x30d,0x30e分别代表红桃2,3,4,5,6,7,8,9,10,J,Q,K,A
// 0x402,0x403,0x404,0x405,0x406,0x407,0x408,0x409,0x40a,0x40b,0x40c,0x40d,0x40e分别代表黑桃2,3,4,5,6,7,8,9,10,J,Q,K,A
// Poker_JokerLittle代表小王
// Poker_JokerBig代表大王
// 小王大王可以变为任意牌，要求算出小王大王变换后最大牌型
// 返回的数字含义如下：
// 1、皇家同花顺：如果花色一样，数字分别是10,J,Q,K,A
// 2、同花顺：如果花色一样，数字是连续的，皇家同花顺除外
// 3、金刚：其中4张牌数字一样
// 4、葫芦：其中3张牌数字一样，另外2张牌数字一样
// 5、同花：花色一样，数字不连续
// 6、顺子：数字是连续，花色不一样
// 7、三条：其中3张牌数字一样，另外2张牌数字不一样
// 8、两对：其中2张牌数字一样，另外其中2张牌数字一样，最后一张数字不一样
// 9、一对：其中2张牌数字一样，另外数字不一样
// 10、高牌：什么都不是
type PokerChanger func([]int, int) bool

var changers [11]PokerChanger = [11]PokerChanger{}

func init() {
	changers[PokerHands_RoyalFlush] = func(input []int, ghost int) bool {
		sort.Ints(input)

		checkNum := len(input) - ghost
		if !isFlush(input[:checkNum]) {
			return false
		}

		match := 0
		for _, card := range input[:checkNum] {
			number := card & Mask_PokerNumber
			if number >= 0x00a && number <= 0x00e {
				match++
			}
		}

		return match == checkNum
	}

	changers[PokerHands_StraightFlush] = func(input []int, ghost int) bool {
		sort.Ints(input)

		checkNum := len(input) - ghost
		return isFlush(input[:checkNum]) && isStraight(input[:checkNum])
	}

	changers[PokerHands_FourOfTheKind] = func(input []int, ghost int) bool {
		sort.Ints(input)

		var (
			temp     [15]int
			checkNum = len(input) - ghost
		)

		for _, card := range input[:checkNum] {
			number := card & Mask_PokerNumber
			temp[number]++
		}

		for _, num := range temp {
			if num >= 3 {
				return true
			}

			if num == 2 && ghost == 2 {
				return true
			}
		}

		return false
	}

	changers[PokerHands_FullHouse] = func(input []int, ghost int) bool {
		sort.Ints(input)

		var (
			temp     [15]int
			checkNum = len(input) - ghost
		)

		for _, card := range input[:checkNum] {
			number := card & Mask_PokerNumber
			temp[number]++
		}

		var hasPair bool
		for _, num := range temp {
			if num >= 2 {
				if ghost == 2 || hasPair {
					return true
				}
				hasPair = true
			}
		}

		return false
	}

	changers[PokerHands_Flush] = func(input []int, ghost int) bool {
		checkNum := len(input) - ghost
		return isFlush(input[:checkNum])
	}

	changers[PokerHands_Straight] = func(input []int, ghost int) bool {
		checkNum := len(input) - ghost
		return isStraight(input[:checkNum])
	}

	changers[PokerHands_ThreeOfAKind] = func(input []int, ghost int) bool {
		if ghost == 2 {
			return true
		}

		var (
			temp     [15]int
			checkNum = len(input) - ghost
		)

		for _, card := range input[:checkNum] {
			number := card & Mask_PokerNumber
			temp[number]++
		}

		for _, num := range temp {
			if num >= 2 {
				return true
			}
		}

		return false
	}

	changers[PokerHands_TwoPairs] = func(input []int, ghost int) bool {
		return changers[PokerHands_ThreeOfAKind](input, ghost)
	}

	changers[PokerHands_Pair] = func(input []int, ghost int) bool { return true }
	changers[PokerHands_HighCard] = func(input []int, ghost int) bool { return true }
}

func Poker02(input []int) int {
	var ghost int
	for _, card := range input {
		if card == Poker_JokerLittle || card == Poker_JokerBig {
			ghost++
		}
	}

	if ghost > 0 {
		for kind := 1; kind < len(changers); kind++ {
			if changers[kind](input, ghost) {
				return kind
			}
		}
	}

	if kind, ok := isFlushKind(input); ok {
		return kind
	}

	if kind, ok := isPairsKind(input); ok {
		return kind
	}

	return PokerHands_HighCard
}

func poker02(input []int) int {
	var (
		jockers []int
		pokers  []int
		numbers []int
	)

	for _, card := range input {
		if card == Poker_JokerLittle || card == Poker_JokerBig {
			jockers = append(jockers, card)
		} else {
			pokers = append(pokers, card)
			numbers = append(numbers, card&Mask_PokerNumber)
		}
	}

	sort.Ints(numbers)

	// 皇家同花顺
	isSameFlower := true
	suit := pokers[0] & Mask_PokerSuit
	for i := 1; i < len(pokers); i++ {
		if (pokers[i] & Mask_PokerSuit) != suit {
			isSameFlower = false
			break
		}
	}

	var (
		flagTJQKA int
		numTJQKA  int
		isTJQKA   bool
	)
	for _, card := range numbers {
		if card >= 0x00a && card <= 0x00e {
			flagTJQKA |= 1 << (card - 0x0a)
		}
	}
	for flagTJQKA > 0 {
		if flagTJQKA&0x1 == 0x1 {
			numTJQKA++
		}
		flagTJQKA >>= 1
	}

	isTJQKA = (len(jockers) + numTJQKA) == 5

	if isSameFlower && isTJQKA {
		return PokerHands_RoyalFlush
	}

	// 同花顺
	isSunzi := false
	if numbers[len(numbers)-1] == 0xe {
		betweenTwoAndFive := true
		betweenTenAndK := true
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] >= 0x2 && numbers[i] <= 0x5 {
				if i == 0 || i > 0 && numbers[i] != numbers[i-1] {
					continue
				}
			}
			betweenTwoAndFive = false
		}
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] >= 0xa && numbers[i] <= 0xd {
				if i == 0 || i > 0 && numbers[i] != numbers[i-1] {
					continue
				}
			}
			betweenTenAndK = false
		}
		isSunzi = betweenTwoAndFive || betweenTenAndK
	} else {
		for i := 1; i < len(numbers); i++ {
			if numbers[i]-numbers[0] < 5 {
				if i > 0 && numbers[i] != numbers[i-1] {
					isSunzi = true
					continue
				}
			}
			isSunzi = false
			break
		}
	}

	if isSameFlower && isSunzi {
		return PokerHands_StraightFlush
	}

	// 四条
	var temp [15]int
	for _, card := range numbers {
		temp[card]++
	}

	for _, num := range temp {
		if len(jockers)+num >= 4 {
			return PokerHands_FourOfTheKind
		}
	}

	// 三带二
	var (
		hasThree bool
		hasPair1 bool
		hasPair2 bool
	)

	for _, num := range temp {
		if num >= 3 {
			hasThree = true
		} else if num == 2 {
			if !hasPair1 {
				hasPair1 = true
			} else {
				hasPair2 = true
			}
		}
	}

	if hasThree || hasPair1 && len(jockers) == 2 || hasPair1 && hasPair2 {
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
	if hasThree || hasPair1 || len(jockers) == 2 {
		return PokerHands_ThreeOfAKind
	}

	// 两对
	if hasPair1 && hasPair2 || hasPair1 && len(jockers) == 0 || len(jockers) == 2 {
		return PokerHands_TwoPairs
	}

	// 一对
	if len(jockers) == 1 {
		return PokerHands_Pair
	}

	return PokerHands_HighCard
}

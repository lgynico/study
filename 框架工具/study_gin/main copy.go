package main

const (
	Mask_PockerSuit   = 0xF00
	Mask_PockerNumber = 0xF
)

const (
	PockerSuit_Diamond = 0x100
	PockerSuit_Club    = 0x200
	PockerSuit_Heart   = 0x300
	PockerSuit_Spade   = 0x400
)

const (
	PockerHands_RoyalFlush    = 1
	PockerHands_StraightFlush = 2
	PockerHands_FourOfTheKind = 3
	PockerHands_FullHouse     = 4
	PockerHands_Flush         = 5
	PockerHands_Straight      = 6
	PockerHands_ThreeOfAKind  = 7
	PockerHands_TwoPairs      = 8
	PockerHands_Pair          = 9
	PockerHands_Nothing       = 10
)

type (
	PockerChecker     func(cards []int) bool
	PockerCheckerList []func(cards []int) bool
)

var (
	RoyalFlush = func(cards []int) bool {
		var (
			isFlush = true
			isABCDE = false
			suit    = cards[0] & Mask_PockerSuit
		)

		for _, card := range cards {
			if card&Mask_PockerSuit != suit {
				isFlush = false
				break
			}
		}

		for _, card := range cards {
			if card&Mask_PockerSuit != suit {
				isFlush = false
				break
			}
		}

		var total int
		for _, card := range cards {
			total += card & Mask_PockerNumber
		}

		return isFlush && isABCDE
	}
)

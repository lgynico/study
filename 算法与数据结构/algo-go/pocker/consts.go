package poker

const (
	PokerHands_RoyalFlush    = 1
	PokerHands_StraightFlush = 2
	PokerHands_FourOfTheKind = 3
	PokerHands_FullHouse     = 4
	PokerHands_Flush         = 5
	PokerHands_Straight      = 6
	PokerHands_ThreeOfAKind  = 7
	PokerHands_TwoPairs      = 8
	PokerHands_Pair          = 9
	PokerHands_HighCard      = 10
)

const (
	Mask_PokerSuit   = 0xF00
	Mask_PokerNumber = 0xF
)

const (
	PokerSuit_Diamond = 0x100
	PokerSuit_Club    = 0x200
	PokerSuit_Heart   = 0x300
	PokerSuit_Spade   = 0x400
)

const (
	Poker_JokerLittle = 0x50f
	Poker_JokerBig    = 0x610
)

func PokerHandsString(hand int) string {
	switch hand {
	case PokerHands_RoyalFlush:
		return "皇家同花顺"
	case PokerHands_StraightFlush:
		return "同花顺"
	case PokerHands_FourOfTheKind:
		return "四条"
	case PokerHands_FullHouse:
		return "三带二"
	case PokerHands_Flush:
		return "同花"
	case PokerHands_Straight:
		return "顺子"
	case PokerHands_ThreeOfAKind:
		return "三条"
	case PokerHands_TwoPairs:
		return "两对"
	case PokerHands_Pair:
		return "一对"
	}

	return "高牌"
}

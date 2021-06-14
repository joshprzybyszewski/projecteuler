package poker

import "github.com/joshprzybyszewski/cribbage/model"

type HandRank int

// The cards are valued in the order:
// 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace.

const (
	unknownRank   HandRank = 0
	HighCard      HandRank = 1  // Highest value card.
	OnePair       HandRank = 2  // Two cards of the same value.
	TwoPair       HandRank = 3  // Two different pairs.
	ThreeOfAKind  HandRank = 4  // Three cards of the same value.
	Straight      HandRank = 5  // All cards are consecutive values.
	Flush         HandRank = 6  // All cards of the same suit.
	FullHouse     HandRank = 7  // Three of a kind and a pair.
	FourOfAKind   HandRank = 8  // Four cards of the same value.
	StraightFlush HandRank = 9  // All cards are consecutive values of same suit.
	RoyalFlush    HandRank = 10 // Ten, Jack, Queen, King, Ace, in same suit.
)

func (hr HandRank) String() string {
	switch hr {
	case HighCard:
		return `HighCard`
	case OnePair:
		return `OnePair`
	case TwoPair:
		return `TwoPair`
	case ThreeOfAKind:
		return `ThreeOfAKind`
	case Straight:
		return `Straight`
	case Flush:
		return `Flush`
	case FullHouse:
		return `FullHouse`
	case FourOfAKind:
		return `FourOfAKind`
	case StraightFlush:
		return `StraightFlush`
	case RoyalFlush:
		return `RoyalFlush`
	default:
		return `unknown`
	}
}

func getRank(sortedByValue []model.Card) HandRank {
	straight := isStraight(sortedByValue)
	flush := isFlush(sortedByValue)
	alike1, alike2 := numAlike(sortedByValue)

	switch {
	case straight && flush && sortedByValue[0].Value == 10:
		return RoyalFlush
	case straight && flush:
		return StraightFlush
	case alike1 == 4, alike2 == 4:
		return FourOfAKind
	case alike1+alike2 == 5:
		return FullHouse
	case flush:
		return Flush
	case straight:
		return Straight
	case alike1 == 3, alike2 == 3:
		return ThreeOfAKind
	case alike1 == 2 && alike2 == 2:
		return TwoPair
	case alike1 == 2, alike2 == 2:
		return OnePair
	default:
		return HighCard
	}
}

func isStraight(sortedByValue []model.Card) bool {
	for i := 1; i < len(sortedByValue)-1; i++ {
		if (sortedByValue[i].Value == 1 && sortedByValue[i-1].Value != 13) ||
			sortedByValue[i].Value != sortedByValue[i-1].Value+1 {
			return false
		}
	}

	return true
}

func isFlush(sortedByValue []model.Card) bool {
	s := sortedByValue[0].Suit
	for i := 1; i < len(sortedByValue)-1; i++ {
		if sortedByValue[i].Suit != s {
			return false
		}
	}

	return true
}

func numAlike(sortedByValue []model.Card) (int, int) {
	first, second := 0, 0
	next := false
	for i := 1; i < len(sortedByValue); i++ {
		if sortedByValue[i].Value == sortedByValue[i-1].Value {
			if next {
				second++
			} else {
				first++
			}
		} else {
			if next && second > 0 {
				break
			}
			if first > 0 {
				next = true
			}
		}
	}
	if first > 0 {
		first++
	}
	if second > 0 {
		second++
	}

	return first, second
}

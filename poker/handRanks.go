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
	alike1, alike2 := getAlike(sortedByValue)

	switch {
	case straight && flush && sortedByValue[0].Value == 1:
		return RoyalFlush
	case straight && flush:
		return StraightFlush
	case len(alike1) == 4, len(alike2) == 4:
		return FourOfAKind
	case len(alike1)+len(alike2) == 5:
		return FullHouse
	case flush:
		return Flush
	case straight:
		return Straight
	case len(alike1) == 3, len(alike2) == 3:
		return ThreeOfAKind
	case len(alike1) == 2 && len(alike2) == 2:
		return TwoPair
	case len(alike1) == 2, len(alike2) == 2:
		return OnePair
	default:
		return HighCard
	}
}

func isStraight(sortedByValue []model.Card) bool {
	for i := 1; i < len(sortedByValue); i++ {
		if i == 1 && sortedByValue[i].Value == 13 {
			if sortedByValue[0].Value != 1 {
				return false
			}
		} else if sortedByValue[i].Value+1 != sortedByValue[i-1].Value {
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

func getAlike(sortedByValue []model.Card) ([]model.Card, []model.Card) {
	var first, second []model.Card
	next := false
	for i := 1; i < len(sortedByValue); i++ {
		if sortedByValue[i].Value == sortedByValue[i-1].Value {
			if next {
				if len(second) == 0 {
					second = append(second, sortedByValue[i-1])
				}
				second = append(second, sortedByValue[i])
			} else {
				if len(first) == 0 {
					first = append(first, sortedByValue[i-1])
				}
				first = append(first, sortedByValue[i])
			}
		} else {
			if next && len(second) > 0 {
				break
			}
			if len(first) > 0 {
				next = true
			}
		}
	}

	return first, second
}

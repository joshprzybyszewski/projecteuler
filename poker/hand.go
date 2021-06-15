package poker

import (
	"fmt"
	"sort"
	"strings"

	"github.com/joshprzybyszewski/cribbage/model"
)

type Hand struct {
	Cards []model.Card
	rank  HandRank
}

func NewHand(cards []string) Hand {
	c := make([]model.Card, len(cards))
	for i, card := range cards {
		card = strings.Replace(card, `T`, `10`, 1)
		c[i] = model.NewCardFromString(card)
	}
	sort.Slice(c, func(i, j int) bool {
		return stableSubCards(c[i], c[j]) > 0
	})

	return Hand{
		Cards: c,
		rank:  getRank(c),
	}
}

func stableSubCards(a, b model.Card) int {
	s := subCards(a, b)
	if s == 0 {
		return int(b.Suit - a.Suit)
	}
	return s
}

// subCards return a-b
// if b is valued greater than a, then returns < 0
// if b is valued less than a, then returns > 0
// if a == b, returns 0
func subCards(a, b model.Card) int {
	if a.Value == b.Value {
		return 0
	}
	if a.Value == 1 {
		return 1
	}
	if b.Value == 1 {
		return -1
	}
	return a.Value - b.Value
}

func (h Hand) Rank() HandRank {
	return h.rank
}

func (h Hand) String() string {
	return fmt.Sprintf("{%s, %v}", h.rank, h.Cards)
}

func (h Hand) Beats(other Hand) bool {
	if h.Rank() == other.Rank() {
		// If two players have the same ranked hands then the
		// rank made up of the highest value wins; for example,
		// a pair of eights beats a pair of fives (see example
		// 1 below). But if two ranks tie, for example, both
		// players have a pair of queens, then highest cards
		// in each hand are compared (see example 4 below);
		// if the highest cards tie then the next highest cards
		// are compared, and so on.

		ma1, ma2 := getAlike(h.Cards)
		oa1, oa2 := getAlike(other.Cards)
		switch h.Rank() {
		case FourOfAKind:
			return subCards(ma1[0], oa1[0]) > 0
		case FullHouse:
			myThree := ma1
			if len(myThree) != 3 {
				myThree = ma2
			}
			otherThree := oa1
			if len(otherThree) != 3 {
				otherThree = oa2
			}
			return subCards(myThree[0], otherThree[0]) > 0
		case ThreeOfAKind, OnePair:
			d := subCards(ma1[0], oa1[0])
			if d != 0 {
				return d > 0
			}
		case TwoPair:
			mbetter := ma1[0]
			mother := ma2[0]
			if subCards(mbetter, mother) < 0 {
				mbetter, mother = mother, mbetter
			}

			obetter := oa1[0]
			oother := oa2[0]
			if subCards(obetter, oother) < 0 {
				obetter, oother = oother, obetter
			}

			d := subCards(mbetter, obetter)
			if d != 0 {
				return d > 0
			}
			d = subCards(mother, oother)
			if d != 0 {
				return d > 0
			}
		}

		for i := 0; i < len(h.Cards); i++ {
			d := subCards(h.Cards[i], other.Cards[i])
			if d != 0 {
				return d > 0
			}
		}
		panic(`impossible!`)
	}
	return h.Rank() > other.Rank()
}

package poker

import (
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
		if c[i].Value != c[j].Value {
			if c[i].Value == 1 {
				return false
			}
			return c[i].Value < c[j].Value
		}
		return c[i].Suit < c[j].Suit
	})

	return Hand{
		Cards: c,
		rank:  getRank(c),
	}
}

func (h Hand) Rank() HandRank {
	return h.rank
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
		//TODO
		panic(`TODO`)
	}
	return h.Rank() > other.Rank()
}

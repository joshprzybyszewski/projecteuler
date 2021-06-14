package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHand(t *testing.T) {
	testCases := []struct {
		cards   []string
		expRank HandRank
	}{{
		cards:   []string{`qh`, `jh`, `kh`, `10h`, `ah`},
		expRank: RoyalFlush,
	}, {
		cards:   []string{`qh`, `jh`, `kh`, `10h`, `9h`},
		expRank: StraightFlush,
	}, {
		cards:   []string{`9s`, `9c`, `kh`, `9d`, `9h`},
		expRank: FourOfAKind,
	}, {
		cards:   []string{`9s`, `7c`, `7h`, `9d`, `9h`},
		expRank: FullHouse,
	}, {
		cards:   []string{`9s`, `7s`, `2s`, `8s`, `ks`},
		expRank: Flush,
	}, {
		cards:   []string{`9s`, `7d`, `6c`, `8s`, `10s`},
		expRank: Straight,
	}, {
		cards:   []string{`9s`, `9d`, `9c`, `8s`, `10s`},
		expRank: ThreeOfAKind,
	}, {
		cards:   []string{`9s`, `9d`, `10c`, `8s`, `10s`},
		expRank: TwoPair,
	}, {
		cards:   []string{`2s`, `9d`, `10c`, `8s`, `10s`},
		expRank: OnePair,
	}, {
		cards:   []string{`2s`, `9d`, `jc`, `8s`, `10s`},
		expRank: HighCard,
	}}

	for _, tc := range testCases {
		h := NewHand(tc.cards)
		assert.Equal(t, tc.expRank, h.Rank(), `expected %s from %+v, but was %s`, tc.expRank, h.Cards, h.Rank())
	}
}

package poker

import (
	"strings"
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

func TestHandBeats(t *testing.T) {
	testCases := []struct {
		both string
		exp  bool
	}{{
		both: `5H 5C 6S 7S KD 2C 3S 8S 8D TD`,
		exp:  false,
	}, {
		both: `5D 8C 9S JS AC 2C 5C 7D 8S QH`,
		exp:  true,
	}, {
		both: `2D 9C AS AH AC 3D 6D 7D TD QD`,
		exp:  false,
	}, {
		both: `4D 6S 9H QH QC 3D 6D 7H QD QS`,
		exp:  true,
	}, {
		both: `2H 2D 4C 4D 4S 3C 3D 3S 9S 9D`,
		exp:  true,
	}}

	for _, tc := range testCases {
		all := strings.Split(tc.both, ` `)
		h1 := NewHand(all[:5])
		h2 := NewHand(all[5:])

		act := h1.Beats(h2)
		if tc.exp {
			assert.True(t, act, `expected %s to beat %s`, h1, h2)
		} else {
			assert.False(t, act, `expected %s to beat %s`, h2, h1)
		}
		opp := h2.Beats(h1)
		assert.NotEqual(t, act, opp)

	}
}

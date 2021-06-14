package easy

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/joshprzybyszewski/projecteuler/poker"
)

const (
	problem54PokerHands = `easy/p054_poker.txt`
)

func SolveProblem54() string {
	/*
		In the card game poker, a hand consists of five cards and
		are ranked, from lowest to highest, in the following way:

		High Card: Highest value card.
		One Pair: Two cards of the same value.
		Two Pairs: Two different pairs.
		Three of a Kind: Three cards of the same value.
		Straight: All cards are consecutive values.
		Flush: All cards of the same suit.
		Full House: Three of a kind and a pair.
		Four of a Kind: Four cards of the same value.
		Straight Flush: All cards are consecutive values of same suit.
		Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.

		The cards are valued in the order:
		2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace.

		The file, poker.txt, contains one-thousand random hands dealt
		to two players. Each line of the file contains ten cards
		(separated by a single space): the first five are Player 1's
		cards and the last five are Player 2's cards. You can assume
		that all hands are valid (no invalid characters or repeated
		cards), each player's hand is in no specific order, and in
		each hand there is a clear winner.

		How many hands does Player 1 win?
	*/

	ans := getProblem54Answer()
	return fmt.Sprintf("%v", ans)
}

func getProblem54Answer() int {
	hands := getProbelm54Cards()

	sum := 0
	for _, bothHands := range hands {
		if len(bothHands) != 10 {
			continue
		}
		if isPlayerOneWin(bothHands[:5], bothHands[5:]) {
			sum++
		}
	}
	return sum
}

func getProbelm54Cards() [][]string {
	dat, err := ioutil.ReadFile(problem54PokerHands)
	if err != nil {
		log.Fatalf("problem 54 cannot read file: %v", err)
	}
	lines := strings.Split(string(dat), "\n")

	res := make([][]string, len(lines))
	for i, line := range lines {
		res[i] = strings.Split(line, ` `)
	}

	return res
}

func isPlayerOneWin(one, two []string) bool {
	h1 := poker.NewHand(one)
	h2 := poker.NewHand(two)

	return h1.Beats(h2)
}

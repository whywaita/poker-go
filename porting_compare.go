package poker

import (
	"fmt"
)

// CompareVSMadeHand compares the hand of a player against all possible boards.
func CompareVSMadeHand(p1 Player) error {
	deck := NewDeck()
	for _, c := range p1.Hand {
		deck.removeCard(c)
	}

	for _, board := range AllCombinations(deck.Cards, 5) {
		madehand := NewBestMadeHand(append(p1.Hand, board...))

		score, hand, err := p1.Evaluate(board)
		if err != nil {
			return err
		}

		if madehand.Type() != score {
			fmt.Println("madehand:", madehand.Type(), "score:", score, "hand:", hand, "board:", board, "p1:", p1.Hand)
		}
	}

	return nil
}

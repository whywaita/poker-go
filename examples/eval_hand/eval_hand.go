package main

import (
	"fmt"
	"log"

	"github.com/whywaita/poker-go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	h1 := []poker.Card{
		{Rank: poker.RankAce, Suit: poker.Hearts},
		{Rank: poker.RankAce, Suit: poker.Diamonds},
	}
	p1 := poker.NewPlayer("player1", h1)

	board := []poker.Card{
		{Rank: poker.RankAce, Suit: poker.Clubs},
		{Rank: poker.RankAce, Suit: poker.Spades},
		{Rank: poker.RankKing, Suit: poker.Hearts},
		{Rank: poker.RankQueen, Suit: poker.Hearts},
		{Rank: poker.RankJack, Suit: poker.Hearts},
	}

	handtype, cards, err := poker.Evaluate(append(p1.Hand, board...))
	if err != nil {
		return fmt.Errorf("poker.Evaluate(append(%s. %s...)): %w", p1.Hand, board, err)
	}

	fmt.Printf("handtype: %s, cards: %s", handtype, cards)

	return nil
}

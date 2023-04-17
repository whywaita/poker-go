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
		{Rank: poker.RankDeuce, Suit: poker.Hearts},
		{Rank: poker.RankThree, Suit: poker.Diamonds},
	}
	p1 := poker.NewPlayer("player1", h1)

	h2 := []poker.Card{
		{Rank: poker.RankAce, Suit: poker.Hearts},
		{Rank: poker.RankAce, Suit: poker.Diamonds},
	}
	p2 := poker.NewPlayer("player2", h2)

	h3 := []poker.Card{
		{Rank: poker.RankSeven, Suit: poker.Clubs},
		{Rank: poker.RankEight, Suit: poker.Clubs},
	}
	p3 := poker.NewPlayer("player3", h3)

	equities, err := poker.EvaluateEquity([]poker.Player{*p1, *p2, *p3})
	if err != nil {
		return fmt.Errorf("failed to evaluate equity: %w", err)
	}
	fmt.Printf("%v equity: %f, %v equity: %f, %v equity %f\n", h1, equities[0], h2, equities[1], h3, equities[2])

	return nil
}

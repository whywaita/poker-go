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
	//winner, err := poker.CompareHands(h1, h2, []poker.Card{
	//	{Rank: 6, Suit: poker.Clubs},
	//	{Rank: 7, Suit: poker.Clubs},
	//	{Rank: 8, Suit: poker.Diamonds},
	//	{Rank: 9, Suit: poker.Clubs},
	//	{Rank: 10, Suit: poker.Clubs},
	//})
	//if err != nil {
	//	return err
	//}
	//fmt.Println(winner)

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
	h1Equity, h2Equity, err := poker.EvaluateEquity(*p1, *p2)
	if err != nil {
		return fmt.Errorf("failed to evaluate equity: %w", err)
	}
	fmt.Printf("%v equity: %f, %v equity: %f\n", h1, h1Equity, h2, h2Equity)

	//if err := poker.CompareVSMadeHand(poker.Player{
	//	Name: "play",
	//	Hand: h2,
	//}); err != nil {
	//	return fmt.Errorf("failed to compare vs made hand: %w", err)
	//}

	return nil
}

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
	//h, err := poker.Evaluate(
	//	[]poker.Card{
	//		{Rank: 8, Suit: poker.Spades},
	//		{Rank: 8, Suit: poker.Diamonds},
	//		{Rank: 8, Suit: poker.Hearts},
	//		{Rank: 3, Suit: poker.Spades},
	//		{Rank: 3, Suit: poker.Hearts},
	//		{Rank: 14, Suit: poker.Clubs},
	//		{Rank: 14, Suit: poker.Diamonds},
	//	})
	//if err != nil {
	//	return fmt.Errorf("Evaluate(): %w", err)
	//}
	//fmt.Println(h)

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
		{Rank: 2, Suit: poker.Hearts},
		{Rank: 3, Suit: poker.Diamonds},
	}

	h2 := []poker.Card{
		{Rank: 14, Suit: poker.Hearts},
		{Rank: 14, Suit: poker.Diamonds},
	}
	h1Equity, h2Equity, err := poker.EvaluateEquity(h1, h2)
	if err != nil {
		return fmt.Errorf("failed to evaluate equity: %w", err)
	}
	fmt.Printf("%v equity: %f, %v equity: %f\n", h1, h1Equity, h2, h2Equity)

	return nil
}

package benchmarks

import (
	"testing"

	"github.com/whywaita/poker-go"
)

func BenchmarkCalcEquity(b *testing.B) {
	in := []poker.Player{
		{
			Name: "player1",
			Hand: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
			},
		},
		{
			Name: "player2",
			Hand: []poker.Card{
				{Rank: poker.RankAce, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Diamonds},
			},
		},
	}

	b.Run("poker.EvaluateEquity", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := poker.EvaluateEquity(in)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("poker.EvaluateEquityByMadeHand", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := poker.EvaluateEquityByMadeHand(in)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

}

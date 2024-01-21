package poker_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/whywaita/poker-go"
)

func TestCompareHandsByMadeHand(t *testing.T) {
	type input struct {
		players []poker.Player
		board   []poker.Card
	}

	tests := []struct {
		name  string
		input input
		want  []poker.Player
	}{
		{
			name: "AA is strong",
			input: input{
				players: []poker.Player{
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
				},
				board: []poker.Card{
					{Rank: poker.RankAce, Suit: poker.Spades},
					{Rank: poker.RankAce, Suit: poker.Clubs},
					{Rank: poker.RankKing, Suit: poker.Clubs},
					{Rank: poker.RankKing, Suit: poker.Diamonds},
					{Rank: poker.RankKing, Suit: poker.Hearts},
				},
			},
			want: []poker.Player{
				{
					Name: "player2",
					Hand: []poker.Card{
						{Rank: poker.RankAce, Suit: poker.Hearts},
						{Rank: poker.RankAce, Suit: poker.Diamonds},
					},
					Score: poker.HandTypeFourOfAKind,
				},
			},
		},
		{
			name: "AA is not always winner",
			input: input{
				players: []poker.Player{
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
				},
				board: []poker.Card{
					{Rank: poker.RankAce, Suit: poker.Spades},
					{Rank: poker.RankFour, Suit: poker.Clubs},
					{Rank: poker.RankFive, Suit: poker.Clubs},
					{Rank: poker.RankNine, Suit: poker.Diamonds},
					{Rank: poker.RankTen, Suit: poker.Hearts},
				},
			},
			want: []poker.Player{
				{
					Name: "player1",
					Hand: []poker.Card{
						{Rank: poker.RankDeuce, Suit: poker.Hearts},
						{Rank: poker.RankThree, Suit: poker.Diamonds},
					},
					Score: poker.HandTypeStraight,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := poker.CompareHandsByMadeHand(tt.input.players, tt.input.board)
			fmt.Println(got)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEvaluateEquityByMadeHand(t *testing.T) {
	tests := []struct {
		name  string
		input []poker.Player
		want  []float64
	}{
		{
			name: "two players",
			input: []poker.Player{
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
			},
			want: []float64{
				0.12031274820359002,
				0.8796866677879629,
			},
		},
		{
			name: "three players",
			input: []poker.Player{
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
				{
					Name: "player3",
					Hand: []poker.Card{
						{Rank: poker.RankEight, Suit: poker.Spades},
						{Rank: poker.RankSeven, Suit: poker.Spades},
					},
				},
			},
			want: []float64{
				0.08533040939512122,
				0.6633188741378833,
				0.2513507164669955,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := poker.EvaluateEquityByMadeHand(tt.input)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEvaluateEquityByMadeHandWithCommunity(t *testing.T) {
	type input struct {
		players   []poker.Player
		community []poker.Card
	}

	tests := []struct {
		name  string
		input input
		want  []float64
	}{
		{
			name: "two players - flop open-end straight draw",
			input: input{
				players: []poker.Player{
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
				},
				community: []poker.Card{
					{Rank: poker.RankFour, Suit: poker.Spades},
					{Rank: poker.RankFive, Suit: poker.Spades},
					{Rank: poker.RankEight, Suit: poker.Spades},
				},
			},
			want: []float64{
				0.23636363636363636,
				0.7636363636363637,
			},
		},
		{
			name: "two players - turn made straight",
			input: input{
				players: []poker.Player{
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
				},
				community: []poker.Card{
					{Rank: poker.RankFour, Suit: poker.Spades},
					{Rank: poker.RankFive, Suit: poker.Spades},
					{Rank: poker.RankEight, Suit: poker.Spades},
					{Rank: poker.RankSix, Suit: poker.Clubs},
				},
			},
			want: []float64{
				0.9545454545454546,
				0.045454545454545456,
			},
		},
		{
			name: "two players - river made straight in community",
			input: input{
				players: []poker.Player{
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
				},
				community: []poker.Card{
					{Rank: poker.RankFour, Suit: poker.Spades},
					{Rank: poker.RankFive, Suit: poker.Spades},
					{Rank: poker.RankEight, Suit: poker.Spades},
					{Rank: poker.RankSix, Suit: poker.Clubs},
					{Rank: poker.RankSeven, Suit: poker.Diamonds},
				},
			},
			want: []float64{
				0,
				0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := poker.EvaluateEquityByMadeHandWithCommunity(tt.input.players, tt.input.community)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

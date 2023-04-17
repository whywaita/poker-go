package poker_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/whywaita/poker-go"
)

func TestIsStraight(t *testing.T) {
	tests := []struct {
		name  string
		input []poker.Card
		want  *[]poker.Card
	}{
		{
			name: "normal straight",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: &[]poker.Card{
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
			},
		},
		{
			name: "normal straight ace to five",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankEight, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: &[]poker.Card{
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
		},
		{
			name: "normal straight ace to ten",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankTen, Suit: poker.Hearts},
				{Rank: poker.RankJack, Suit: poker.Spades},
				{Rank: poker.RankQueen, Suit: poker.Diamonds},
				{Rank: poker.RankKing, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: &[]poker.Card{
				{Rank: poker.RankAce, Suit: poker.Spades},
				{Rank: poker.RankKing, Suit: poker.Hearts},
				{Rank: poker.RankQueen, Suit: poker.Diamonds},
				{Rank: poker.RankJack, Suit: poker.Spades},
				{Rank: poker.RankTen, Suit: poker.Hearts},
			},
		},
		{
			name: "not straight",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFive, Suit: poker.Hearts},
				{Rank: poker.RankSix, Suit: poker.Spades},
				{Rank: poker.RankEight, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := poker.IsStraight(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsStraightFlush(t *testing.T) {
	tests := []struct {
		name  string
		input []poker.Card
		want  []poker.Card
	}{
		{
			name: "normal straight flush",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Clubs},
				{Rank: poker.RankFour, Suit: poker.Clubs},
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankSix, Suit: poker.Clubs},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Hearts},
			},
			want: []poker.Card{
				{Rank: poker.RankSix, Suit: poker.Clubs},
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankFour, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
			},
		},
		{
			name: "normal straight flush ace to five",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Clubs},
				{Rank: poker.RankFour, Suit: poker.Clubs},
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankEight, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Diamonds},
				{Rank: poker.RankAce, Suit: poker.Clubs},
			},
			want: []poker.Card{
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankFour, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankAce, Suit: poker.Clubs},
			},
		},
		{
			name: "straight flush is high than straight",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Clubs},
				{Rank: poker.RankFour, Suit: poker.Clubs},
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankEight, Suit: poker.Diamonds},
				{Rank: poker.RankAce, Suit: poker.Clubs},
			},
			want: []poker.Card{
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankFour, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankAce, Suit: poker.Clubs},
			},
		},
		{
			name: "straight flush is high than straight",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Hearts},
				{Rank: poker.RankSix, Suit: poker.Hearts},
				{Rank: poker.RankSeven, Suit: poker.Hearts},
			},
			want: []poker.Card{
				{Rank: poker.RankSeven, Suit: poker.Hearts},
				{Rank: poker.RankSix, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Hearts},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Hearts},
			},
		},
		{
			name: "not straight flush",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFive, Suit: poker.Hearts},
				{Rank: poker.RankSix, Suit: poker.Spades},
				{Rank: poker.RankEight, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := poker.IsStraightFlush(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

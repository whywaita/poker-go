package poker_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/whywaita/poker-go"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name  string
		input []poker.Card
		want  poker.HandType
	}{
		{
			name: "straight",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeStraight,
		},
		{
			name: "straight is higher than pair",
			input: []poker.Card{
				{Rank: poker.RankSix, Suit: poker.Clubs},
				{Rank: poker.RankSeven, Suit: poker.Diamonds},
				{Rank: poker.RankEight, Suit: poker.Hearts},
				{Rank: poker.RankNine, Suit: poker.Spades},
				{Rank: poker.RankTen, Suit: poker.Diamonds},
				{Rank: poker.RankAce, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeStraight,
		},
		{
			name: "High card",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFive, Suit: poker.Hearts},
				{Rank: poker.RankSix, Suit: poker.Spades},
				{Rank: poker.RankEight, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeHighCard,
		},
		{
			name: "royal flush",
			input: []poker.Card{
				{Rank: poker.RankTen, Suit: poker.Spades},
				{Rank: poker.RankJack, Suit: poker.Spades},
				{Rank: poker.RankQueen, Suit: poker.Spades},
				{Rank: poker.RankKing, Suit: poker.Spades},
				{Rank: poker.RankAce, Suit: poker.Spades},
				{Rank: poker.RankSix, Suit: poker.Clubs},
				{Rank: poker.RankSeven, Suit: poker.Diamonds},
			},
			want: poker.HandTypeRoyalFlush,
		},
		{
			name: "straight flush",
			input: []poker.Card{
				{Rank: poker.RankSeven, Suit: poker.Spades},
				{Rank: poker.RankEight, Suit: poker.Spades},
				{Rank: poker.RankNine, Suit: poker.Spades},
				{Rank: poker.RankTen, Suit: poker.Spades},
				{Rank: poker.RankJack, Suit: poker.Spades},
				{Rank: poker.RankThree, Suit: poker.Hearts},
				{Rank: poker.RankFour, Suit: poker.Diamonds},
			},
			want: poker.HandTypeStraightFlush,
		},
		{
			name: "four of a kind",
			input: []poker.Card{
				{Rank: poker.RankSeven, Suit: poker.Spades},
				{Rank: poker.RankSeven, Suit: poker.Diamonds},
				{Rank: poker.RankSeven, Suit: poker.Hearts},
				{Rank: poker.RankSeven, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Spades},
				{Rank: poker.RankThree, Suit: poker.Clubs},
				{Rank: poker.RankNine, Suit: poker.Hearts},
			},
			want: poker.HandTypeFourOfAKind,
		},
		{
			name: "full house",
			input: []poker.Card{
				{Rank: poker.RankEight, Suit: poker.Spades},
				{Rank: poker.RankEight, Suit: poker.Diamonds},
				{Rank: poker.RankEight, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Spades},
				{Rank: poker.RankThree, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Clubs},
				{Rank: poker.RankAce, Suit: poker.Diamonds},
			},
			want: poker.HandTypeFullHouse,
		},
		{
			name: "flush",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Hearts},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankSix, Suit: poker.Hearts},
				{Rank: poker.RankEight, Suit: poker.Hearts},
				{Rank: poker.RankTen, Suit: poker.Hearts},
				{Rank: poker.RankJack, Suit: poker.Clubs},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeFlush,
		},
		{
			name: "full house",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Diamonds},
				{Rank: poker.RankDeuce, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeFullHouse,
		},
		{
			name: "four of a kind",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Diamonds},
				{Rank: poker.RankDeuce, Suit: poker.Hearts},
				{Rank: poker.RankDeuce, Suit: poker.Spades},
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeFourOfAKind,
		},
		{
			name: "straight flush",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Diamonds},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankFour, Suit: poker.Diamonds},
				{Rank: poker.RankFive, Suit: poker.Diamonds},
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeStraightFlush,
		},
		{
			name: "royal flush",
			input: []poker.Card{
				{Rank: poker.RankTen, Suit: poker.Hearts},
				{Rank: poker.RankJack, Suit: poker.Hearts},
				{Rank: poker.RankQueen, Suit: poker.Hearts},
				{Rank: poker.RankKing, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Hearts},
				{Rank: poker.RankThree, Suit: poker.Diamonds},
				{Rank: poker.RankNine, Suit: poker.Hearts},
			},
			want: poker.HandTypeRoyalFlush,
		},
		{
			name: "three of a kind",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Diamonds},
				{Rank: poker.RankDeuce, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankSeven, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeThreeOfAKind,
		},
		{
			name: "two pair",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Diamonds},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankFour, Suit: poker.Spades},
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankSeven, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypeTwoPair,
		},
		{
			name: "one pair",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Diamonds},
				{Rank: poker.RankFour, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankSix, Suit: poker.Diamonds},
				{Rank: poker.RankSeven, Suit: poker.Hearts},
				{Rank: poker.RankAce, Suit: poker.Spades},
			},
			want: poker.HandTypePair,
		},
		{
			name: "full house - pair and three of a kind",
			input: []poker.Card{
				{Rank: poker.RankDeuce, Suit: poker.Clubs},
				{Rank: poker.RankDeuce, Suit: poker.Diamonds},
				{Rank: poker.RankDeuce, Suit: poker.Hearts},
				{Rank: poker.RankFive, Suit: poker.Spades},
				{Rank: poker.RankFive, Suit: poker.Diamonds},
				{Rank: poker.RankFive, Suit: poker.Clubs},
				{Rank: poker.RankNine, Suit: poker.Spades},
			},
			want: poker.HandTypeFullHouse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := poker.Evaluate(tt.input)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

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

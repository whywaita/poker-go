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
		want  poker.Hand
	}{
		{
			name: "straight",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 6, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandStraight,
		},
		{
			name: "straight is higher than pair",
			input: []poker.Card{
				{Rank: 6, Suit: poker.Clubs},
				{Rank: 7, Suit: poker.Diamonds},
				{Rank: 8, Suit: poker.Hearts},
				{Rank: 9, Suit: poker.Spades},
				{Rank: 10, Suit: poker.Diamonds},
				{Rank: 14, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandStraight,
		},
		{
			name: "High card",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 5, Suit: poker.Hearts},
				{Rank: 6, Suit: poker.Spades},
				{Rank: 8, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandHighCard,
		},
		{
			name: "royal flush",
			input: []poker.Card{
				{Rank: 10, Suit: poker.Spades},
				{Rank: 11, Suit: poker.Spades},
				{Rank: 12, Suit: poker.Spades},
				{Rank: 13, Suit: poker.Spades},
				{Rank: 14, Suit: poker.Spades},
				{Rank: 6, Suit: poker.Clubs},
				{Rank: 7, Suit: poker.Diamonds},
			},
			want: poker.HandRoyalFlush,
		},
		{
			name: "straight flush",
			input: []poker.Card{
				{Rank: 7, Suit: poker.Spades},
				{Rank: 8, Suit: poker.Spades},
				{Rank: 9, Suit: poker.Spades},
				{Rank: 10, Suit: poker.Spades},
				{Rank: 11, Suit: poker.Spades},
				{Rank: 3, Suit: poker.Hearts},
				{Rank: 4, Suit: poker.Diamonds},
			},
			want: poker.HandStraightFlush,
		},
		{
			name: "four of a kind",
			input: []poker.Card{
				{Rank: 7, Suit: poker.Spades},
				{Rank: 7, Suit: poker.Diamonds},
				{Rank: 7, Suit: poker.Hearts},
				{Rank: 7, Suit: poker.Clubs},
				{Rank: 2, Suit: poker.Spades},
				{Rank: 3, Suit: poker.Clubs},
				{Rank: 9, Suit: poker.Hearts},
			},
			want: poker.HandFourOfAKind,
		},
		{
			name: "full house",
			input: []poker.Card{
				{Rank: 8, Suit: poker.Spades},
				{Rank: 8, Suit: poker.Diamonds},
				{Rank: 8, Suit: poker.Hearts},
				{Rank: 3, Suit: poker.Spades},
				{Rank: 3, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Clubs},
				{Rank: 14, Suit: poker.Diamonds},
			},
			want: poker.HandFullHouse,
		},
		{
			name: "flush",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Hearts},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 6, Suit: poker.Hearts},
				{Rank: 8, Suit: poker.Hearts},
				{Rank: 10, Suit: poker.Hearts},
				{Rank: 11, Suit: poker.Clubs},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandFlush,
		},
		{
			name: "full house",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 2, Suit: poker.Diamonds},
				{Rank: 2, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandFullHouse,
		},
		{
			name: "four of a kind",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 2, Suit: poker.Diamonds},
				{Rank: 2, Suit: poker.Hearts},
				{Rank: 2, Suit: poker.Spades},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandFourOfAKind,
		},
		{
			name: "straight flush",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Diamonds},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Diamonds},
				{Rank: 5, Suit: poker.Diamonds},
				{Rank: 6, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandStraightFlush,
		},
		{
			name: "royal flush",
			input: []poker.Card{
				{Rank: 10, Suit: poker.Hearts},
				{Rank: 11, Suit: poker.Hearts},
				{Rank: 12, Suit: poker.Hearts},
				{Rank: 13, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Hearts},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
			},
			want: poker.HandRoyalFlush,
		},
		{
			name: "three of a kind",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 2, Suit: poker.Diamonds},
				{Rank: 2, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 6, Suit: poker.Diamonds},
				{Rank: 7, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandThreeOfAKind,
		},
		{
			name: "two pair",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 2, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 4, Suit: poker.Spades},
				{Rank: 6, Suit: poker.Diamonds},
				{Rank: 7, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandTwoPair,
		},
		{
			name: "one pair",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 2, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 6, Suit: poker.Diamonds},
				{Rank: 7, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: poker.HandPair,
		},
		{
			name: "full house - pair and three of a kind",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 2, Suit: poker.Diamonds},
				{Rank: 2, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 5, Suit: poker.Diamonds},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 9, Suit: poker.Spades},
			},
			want: poker.HandFullHouse,
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
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 6, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: &[]poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 6, Suit: poker.Diamonds},
			},
		},
		{
			name: "normal straight ace to five",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 8, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: &[]poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Spades},
				{Rank: 14, Suit: poker.Spades},
			},
		},
		{
			name: "normal straight ace to ten",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 10, Suit: poker.Hearts},
				{Rank: 11, Suit: poker.Spades},
				{Rank: 12, Suit: poker.Diamonds},
				{Rank: 13, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
			want: &[]poker.Card{
				{Rank: 10, Suit: poker.Hearts},
				{Rank: 11, Suit: poker.Spades},
				{Rank: 12, Suit: poker.Diamonds},
				{Rank: 13, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
			},
		},
		{
			name: "not straight",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 5, Suit: poker.Hearts},
				{Rank: 6, Suit: poker.Spades},
				{Rank: 8, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
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
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Clubs},
				{Rank: 4, Suit: poker.Clubs},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 6, Suit: poker.Clubs},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Hearts},
			},
			want: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Clubs},
				{Rank: 4, Suit: poker.Clubs},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 6, Suit: poker.Clubs},
			},
		},
		{
			name: "normal straight flush ace to five",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Clubs},
				{Rank: 4, Suit: poker.Clubs},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 8, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Diamonds},
				{Rank: 14, Suit: poker.Clubs},
			},
			want: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Clubs},
				{Rank: 4, Suit: poker.Clubs},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 14, Suit: poker.Clubs},
			},
		},
		{
			name: "straight flush is high than straight",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Clubs},
				{Rank: 4, Suit: poker.Clubs},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 6, Suit: poker.Diamonds},
				{Rank: 8, Suit: poker.Diamonds},
				{Rank: 14, Suit: poker.Clubs},
			},
			want: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Clubs},
				{Rank: 4, Suit: poker.Clubs},
				{Rank: 5, Suit: poker.Clubs},
				{Rank: 14, Suit: poker.Clubs},
			},
		},
		{
			name: "straight flush is high than straight",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Hearts},
				{Rank: 3, Suit: poker.Hearts},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Hearts},
				{Rank: 6, Suit: poker.Hearts},
				{Rank: 7, Suit: poker.Hearts},
			},
			want: []poker.Card{
				{Rank: 3, Suit: poker.Hearts},
				{Rank: 4, Suit: poker.Hearts},
				{Rank: 5, Suit: poker.Hearts},
				{Rank: 6, Suit: poker.Hearts},
				{Rank: 7, Suit: poker.Hearts},
			},
		},
		{
			name: "not straight flush",
			input: []poker.Card{
				{Rank: 2, Suit: poker.Clubs},
				{Rank: 3, Suit: poker.Diamonds},
				{Rank: 5, Suit: poker.Hearts},
				{Rank: 6, Suit: poker.Spades},
				{Rank: 8, Suit: poker.Diamonds},
				{Rank: 9, Suit: poker.Hearts},
				{Rank: 14, Suit: poker.Spades},
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

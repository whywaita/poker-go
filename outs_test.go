package poker

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalculateOuts(t *testing.T) {
	tests := []struct {
		name           string
		holeCards      []Card
		board          []Card
		opponents      [][]Card
		wantCount      int
		wantCards      []Card
		wantContains   []Card
		wantNotContain []Card
		wantErr        bool
	}{
		{
			name: "KhQd vs 7h7d board Jc9c8c3s - outs are K Q T only",
			holeCards: []Card{
				{RankKing, Hearts},
				{RankQueen, Diamonds},
			},
			board: []Card{
				{RankJack, Clubs},
				{RankNine, Clubs},
				{RankEight, Clubs},
				{RankThree, Spades},
			},
			opponents: [][]Card{
				{{RankSeven, Hearts}, {RankSeven, Diamonds}},
			},
			wantCount: 10,
			wantContains: []Card{
				{RankKing, Clubs},
				{RankKing, Diamonds},
				{RankKing, Spades},
				{RankQueen, Hearts},
				{RankQueen, Clubs},
				{RankQueen, Spades},
				{RankTen, Hearts},
				{RankTen, Clubs},
				{RankTen, Diamonds},
				{RankTen, Spades},
			},
			wantNotContain: []Card{
				{RankThree, Hearts},
				{RankThree, Clubs},
				{RankThree, Diamonds},
				{RankJack, Diamonds},
				{RankEight, Diamonds},
				{RankNine, Diamonds},
			},
		},
		{
			name: "already winning - no outs",
			holeCards: []Card{
				{RankAce, Hearts},
				{RankKing, Hearts},
			},
			board: []Card{
				{RankAce, Diamonds},
				{RankKing, Diamonds},
				{RankFive, Spades},
				{RankSeven, Clubs},
			},
			opponents: [][]Card{
				{{RankDeuce, Clubs}, {RankThree, Clubs}},
			},
			wantCount: 0,
		},
		{
			name: "flop board=3 - AhKd vs 9h9d board 7c4s2h",
			holeCards: []Card{
				{RankAce, Hearts},
				{RankKing, Diamonds},
			},
			board: []Card{
				{RankSeven, Clubs},
				{RankFour, Spades},
				{RankDeuce, Hearts},
			},
			opponents: [][]Card{
				{{RankNine, Hearts}, {RankNine, Diamonds}},
			},
			wantCount: 6,
			wantContains: []Card{
				{RankAce, Clubs},
				{RankAce, Diamonds},
				{RankAce, Spades},
				{RankKing, Clubs},
				{RankKing, Hearts},
				{RankKing, Spades},
			},
			wantNotContain: []Card{
				{RankSeven, Diamonds},
				{RankFour, Diamonds},
				{RankDeuce, Diamonds},
			},
		},
		{
			name: "river board 5 cards - no outs",
			holeCards: []Card{
				{RankAce, Hearts},
				{RankKing, Hearts},
			},
			board: []Card{
				{RankDeuce, Clubs},
				{RankThree, Diamonds},
				{RankFour, Spades},
				{RankSeven, Hearts},
				{RankNine, Clubs},
			},
			opponents: [][]Card{
				{{RankDeuce, Hearts}, {RankDeuce, Diamonds}},
			},
		},
		{
			name:      "invalid hole cards - 1 card",
			holeCards: []Card{{RankAce, Hearts}},
			board: []Card{
				{RankDeuce, Clubs},
				{RankThree, Diamonds},
				{RankFour, Spades},
			},
			opponents: [][]Card{
				{{RankFive, Hearts}, {RankSix, Hearts}},
			},
			wantErr: true,
		},
		{
			name: "invalid board - 6 cards",
			holeCards: []Card{
				{RankAce, Hearts},
				{RankKing, Hearts},
			},
			board: []Card{
				{RankDeuce, Clubs},
				{RankThree, Diamonds},
				{RankFour, Spades},
				{RankFive, Hearts},
				{RankSix, Clubs},
				{RankSeven, Diamonds},
			},
			opponents: [][]Card{
				{{RankEight, Hearts}, {RankNine, Hearts}},
			},
			wantErr: true,
		},
		{
			name: "invalid board - 2 cards",
			holeCards: []Card{
				{RankAce, Hearts},
				{RankKing, Hearts},
			},
			board: []Card{
				{RankDeuce, Clubs},
				{RankThree, Diamonds},
			},
			opponents: [][]Card{
				{{RankFour, Hearts}, {RankFive, Hearts}},
			},
			wantErr: true,
		},
		{
			name: "invalid opponent - 1 card",
			holeCards: []Card{
				{RankAce, Hearts},
				{RankKing, Hearts},
			},
			board: []Card{
				{RankDeuce, Clubs},
				{RankThree, Diamonds},
				{RankFour, Spades},
			},
			opponents: [][]Card{
				{{RankFive, Hearts}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateOuts(tt.holeCards, tt.board, tt.opponents)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantCount > 0 && len(got) != tt.wantCount {
				t.Errorf("outs count = %d, want %d", len(got), tt.wantCount)
			}

			if tt.wantCards != nil {
				if diff := cmp.Diff(tt.wantCards, got); diff != "" {
					t.Errorf("outs mismatch (-want +got):\n%s", diff)
				}
			}

			for _, wc := range tt.wantContains {
				if !slices.Contains(got, wc) {
					t.Errorf("expected outs to contain %v, got %v", wc, got)
				}
			}

			for _, nc := range tt.wantNotContain {
				if slices.Contains(got, nc) {
					t.Errorf("expected outs NOT to contain %v", nc)
				}
			}
		})
	}
}

func BenchmarkCalculateOuts(b *testing.B) {
	holeCards := []Card{
		{RankKing, Hearts},
		{RankQueen, Diamonds},
	}
	board := []Card{
		{RankJack, Clubs},
		{RankNine, Clubs},
		{RankEight, Clubs},
		{RankThree, Spades},
	}
	opponents := [][]Card{
		{{RankSeven, Hearts}, {RankSeven, Diamonds}},
	}

	for b.Loop() {
		_, _ = CalculateOuts(holeCards, board, opponents)
	}
}

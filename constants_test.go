package poker_test

import (
	"github.com/whywaita/poker-go"
	"testing"
)

func TestCard_StringShort(t *testing.T) {
	tests := []struct {
		name            string // description of this test case
		c               poker.Card
		wantString      string
		wantStringShort string
	}{
		{"Print ", poker.Card{Rank: poker.RankQueen, Suit: poker.Spades}, "Q spades", "Qs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.String()
			if got != tt.wantString {
				t.Errorf("StringShort() = %v, want %v", got, tt.wantString)
			}
			gotShort := tt.c.StringShort()
			if gotShort != tt.wantStringShort {
				t.Errorf("StringShort() = %v, want %v", gotShort, tt.wantStringShort)
			}
		})
	}
}

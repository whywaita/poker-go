package poker

import (
	"fmt"
)

// made_hand.go is porting from github.com/axross/poker

type MadeHand struct {
	Value int
}

func NewMadeHandFromIndex(value int) *MadeHand {
	return &MadeHand{
		Value: value,
	}
}

func NewBestMadeHand(cards []Card) *MadeHand {
	flushSuit := findFlushSuit(cards)

	if flushSuit != nil {
		return NewMadeHandFromIndex(asFlush[hashForFlush(cards, *flushSuit)])
	}

	return NewMadeHandFromIndex(asRainbow[hashForRainbow(cards)])
}

func (hand *MadeHand) Type() HandType {
	if hand.Value > 6185 {
		return HandTypeHighCard
	}
	if hand.Value > 3325 {
		return HandTypePair
	}
	if hand.Value > 2467 {
		return HandTypeTwoPair
	}
	if hand.Value > 1609 {
		return HandTypeThreeOfAKind
	}
	if hand.Value > 1599 {
		return HandTypeStraight
	}
	if hand.Value > 322 {
		return HandTypeFlush
	}
	if hand.Value > 166 {
		return HandTypeFullHouse
	}
	if hand.Value > 10 {
		return HandTypeFourOfAKind
	}

	if hand.Value == 1 {
		return HandTypeRoyalFlush
	}

	return HandTypeStraightFlush

}

func (hand *MadeHand) Power() int {
	return 7462 - hand.Value
}

func (hand *MadeHand) String() string {
	return fmt.Sprintf("MadeHand<%d>", hand.Value)
}

func findFlushSuit(cards []Card) *Suit {
	flush := make(map[Suit]int, 0)

	for _, card := range cards {
		suit, ok := flush[card.Suit]
		if !ok {
			flush[card.Suit] = 1
			continue
		}
		flush[card.Suit] = suit + 1
	}

	for suit, count := range flush {
		if count >= 5 {
			return &suit
		}
	}
	return nil
}

func hashForFlush(cards []Card, flushSuit Suit) int {
	bitEachRank := map[Rank]int{
		RankAce:   0x1000,
		RankDeuce: 0x1,
		RankThree: 0x2,
		RankFour:  0x4,
		RankFive:  0x8,
		RankSix:   0x10,
		RankSeven: 0x20,
		RankEight: 0x40,
		RankNine:  0x80,
		RankTen:   0x100,
		RankJack:  0x200,
		RankQueen: 0x400,
		RankKing:  0x800,
	}
	var hash int

	for _, card := range cards {
		if card.Suit == flushSuit {
			hash += bitEachRank[card.Rank]
		}
	}

	return hash
}

func hashForRainbow(cards []Card) int {
	cardLengthEachRank := map[Rank]int{
		RankDeuce: 0,
		RankThree: 0,
		RankFour:  0,
		RankFive:  0,
		RankSix:   0,
		RankSeven: 0,
		RankEight: 0,
		RankNine:  0,
		RankTen:   0,
		RankJack:  0,
		RankQueen: 0,
		RankKing:  0,
		RankAce:   0,
	}
	remainingCardLength := len(cards)

	for _, card := range cards {
		cardLengthEachRank[card.Rank]++
	}

	hash := 0

	for _, rank := range []Rank{
		RankDeuce,
		RankThree,
		RankFour,
		RankFive,
		RankSix,
		RankSeven,
		RankEight,
		RankNine,
		RankTen,
		RankJack,
		RankQueen,
		RankKing,
		RankAce,
	} {
		length := cardLengthEachRank[rank]

		if length == 0 {
			continue
		}

		hash += dpReference[length][rank][remainingCardLength]
		remainingCardLength -= length

		if remainingCardLength == 0 {
			break
		}
	}

	return hash

}

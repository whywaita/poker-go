package poker

import (
	"fmt"
	"slices"
)

const selfPlayerName = "self"

// CalculateOuts returns the cards from the remaining deck that would
// change the result from self losing to self winning against all opponents.
// holeCards must be exactly 2 cards. board must be 3, 4, or 5 cards.
// Each opponent must have exactly 2 cards.
func CalculateOuts(holeCards []Card, board []Card, opponents [][]Card) ([]Card, error) {
	if len(holeCards) != 2 {
		return nil, fmt.Errorf("holeCards must be exactly 2 cards, got %d", len(holeCards))
	}
	if len(board) < 3 || len(board) > 5 {
		return nil, fmt.Errorf("board must be 3, 4, or 5 cards, got %d", len(board))
	}
	for i, opp := range opponents {
		if len(opp) != 2 {
			return nil, fmt.Errorf("opponent %d must have exactly 2 cards, got %d", i, len(opp))
		}
	}

	if len(board) == 5 {
		return nil, nil
	}

	players := make([]Player, 0, len(opponents)+1)
	players = append(players, Player{Name: selfPlayerName, Hand: holeCards})
	for i, opp := range opponents {
		players = append(players, Player{Name: fmt.Sprintf("opponent%d", i), Hand: opp})
	}

	deck := NewDeck()
	for _, c := range holeCards {
		deck.RemoveCard(c)
	}
	for _, c := range board {
		deck.RemoveCard(c)
	}
	for _, opp := range opponents {
		for _, c := range opp {
			deck.RemoveCard(c)
		}
	}

	beforeWinners, err := compareWinners(players, board)
	if err != nil {
		return nil, fmt.Errorf("failed to compare before: %w", err)
	}
	selfWinsBefore := isPlayerInWinners(selfPlayerName, beforeWinners)

	var outs []Card
	newBoard := make([]Card, len(board)+1)
	copy(newBoard, board)

	for _, card := range deck.Cards {
		newBoard[len(board)] = card

		afterWinners, err := compareWinners(players, newBoard)
		if err != nil {
			return nil, fmt.Errorf("failed to compare after adding %v: %w", card, err)
		}

		if !selfWinsBefore && isPlayerInWinners(selfPlayerName, afterWinners) {
			outs = append(outs, card)
		}
	}

	return outs, nil
}

// compareWinners determines the winner(s) among players given a board.
// Uses CompareHandsByMadeHand for 5-card boards (7 total cards with 2 hole cards),
// falls back to evaluate-based comparison for shorter boards.
func compareWinners(players []Player, board []Card) ([]Player, error) {
	if len(board) == 5 {
		return CompareHandsByMadeHand(players, board)
	}
	return compareHandsByEval(players, board)
}

// compareHandsByEval compares hands using evaluate() for non-7-card totals.
func compareHandsByEval(players []Player, board []Card) ([]Player, error) {
	type eval struct {
		player   Player
		handType HandType
		rankKey  []Rank
	}

	var best *eval
	var tied []Player

	for _, p := range players {
		cards := make([]Card, len(p.Hand)+len(board))
		copy(cards, p.Hand)
		copy(cards[len(p.Hand):], board)

		handType, hand, err := evaluate(cards)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate %s: %w", p.Name, err)
		}

		e := &eval{
			player:   p,
			handType: handType,
			rankKey:  handRankKey(handType, hand),
		}

		if best == nil {
			best = e
			continue
		}

		cmp := compareRankedHands(best.handType, best.rankKey, e.handType, e.rankKey)
		switch {
		case cmp < 0:
			best = e
			tied = nil
		case cmp == 0:
			tied = append(tied, p)
		}
	}

	if len(tied) == 0 {
		return []Player{best.player}, nil
	}
	result := make([]Player, 0, len(tied)+1)
	result = append(result, tied...)
	result = append(result, best.player)
	return result, nil
}

// compareRankedHands compares two hands by HandType and rank key.
// Returns >0 if a is better, <0 if b is better, 0 if tie.
func compareRankedHands(aType HandType, aKey []Rank, bType HandType, bKey []Rank) int {
	if aType != bType {
		if aType > bType {
			return 1
		}
		return -1
	}
	for i := range min(len(aKey), len(bKey)) {
		if aKey[i] > bKey[i] {
			return 1
		}
		if aKey[i] < bKey[i] {
			return -1
		}
	}
	return 0
}

// handRankKey extracts comparison-relevant ranks in priority order.
// The returned slice can be compared lexicographically to determine the stronger hand
// within the same HandType.
func handRankKey(handType HandType, hand []Card) []Rank {
	switch handType {
	case HandTypeRoyalFlush:
		return nil
	case HandTypeStraightFlush:
		// evaluate() returns ascending [low..high]
		if hand[4].Rank == RankAce && hand[0].Rank == RankDeuce {
			return []Rank{RankFive}
		}
		return []Rank{hand[4].Rank}
	case HandTypeFourOfAKind:
		// [Q,Q,Q,Q,kicker]
		return []Rank{hand[0].Rank, hand[4].Rank}
	case HandTypeFullHouse:
		// [T,T,T,P,P]
		return []Rank{hand[0].Rank, hand[3].Rank}
	case HandTypeFlush:
		// descending [high..low]
		return []Rank{hand[0].Rank, hand[1].Rank, hand[2].Rank, hand[3].Rank, hand[4].Rank}
	case HandTypeStraight:
		// descending [high..low]
		if hand[0].Rank == RankFive && hand[4].Rank == RankAce {
			return []Rank{RankFive}
		}
		return []Rank{hand[0].Rank}
	case HandTypeThreeOfAKind:
		// [T,T,T,K1,K2]
		return []Rank{hand[0].Rank, hand[3].Rank, hand[4].Rank}
	case HandTypeTwoPair:
		// [HP,HP,LP,LP,kicker]
		return []Rank{hand[0].Rank, hand[2].Rank, hand[4].Rank}
	case HandTypePair:
		// [P,P,K1,K2,K3]
		return []Rank{hand[0].Rank, hand[2].Rank, hand[3].Rank, hand[4].Rank}
	case HandTypeHighCard:
		// descending [C1,C2,C3,C4,C5]
		return []Rank{hand[0].Rank, hand[1].Rank, hand[2].Rank, hand[3].Rank, hand[4].Rank}
	default:
		return nil
	}
}

func isPlayerInWinners(name string, winners []Player) bool {
	return slices.ContainsFunc(winners, func(p Player) bool {
		return p.Name == name
	})
}

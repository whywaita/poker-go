package poker

import (
	"fmt"
)

// EvaluateEquity returns the equity of each player in the game.
func EvaluateEquity(players []Player) ([]float64, error) {
	deck := NewDeck()

	for _, p := range players {
		for _, c := range p.Hand {
			deck.removeCard(c)
		}
	}

	wins := make([]int, len(players))

	var ties int
	var total int
	for _, board := range AllCombinations(deck.Cards, 5) {
		winners, err := CompareHands(players, board)
		if err != nil {
			return nil, err
		}
		switch {
		case len(winners) == len(players):
			ties++
		default:
			for _, winner := range winners {
				wins[indexOf(winner, players)] += 1
			}
		}
		total++
	}

	equities := make([]float64, 0, len(players))
	baseTies := ties / len(players)

	for _, win := range wins {
		equities = append(equities, float64(win+baseTies)/float64(total))
	}

	return equities, nil
}

// Evaluate returns the best hand type and cards for the given cards.
func Evaluate(cards []Card) (HandType, []Card, error) {
	if len(cards) != 7 {
		return 0, nil, fmt.Errorf("invalid number of cards")
	}
	return evaluate(cards)
}

func evaluate(cards []Card) (HandType, []Card, error) {
	cards = sortByRankDesc(cards)

	flush, _ := isFlush(cards)
	straight := IsStraight(cards)
	straightFlush := IsStraightFlush(cards)
	pairs := GetPairs(cards)
	fullHouse := isFullHouse(pairs)

	if straightFlush != nil {
		sortByRank(straightFlush)
		if straightFlush[4].Rank == RankAce && straightFlush[0].Rank == RankTen {
			return HandTypeRoyalFlush, straightFlush, nil
		}
		return HandTypeStraightFlush, straightFlush, nil
	} else if fullHouse != nil {
		return HandTypeFullHouse, fullHouse, nil
	} else if flush != nil {
		return HandTypeFlush, flush, nil
	} else if straight != nil {
		return HandTypeStraight, *straight, nil
	} else if len(pairs) != 0 {
		switch {
		case len(pairs[0]) == 3:
			threeOfAKind := []Card{pairs[0][0], pairs[0][1], pairs[0][2]}
			for _, card := range pairs[0] {
				cards = removeCards(cards, card)
			}
			sortByRankDesc(cards)
			threeOfAKind = append(threeOfAKind, cards[0])
			threeOfAKind = append(threeOfAKind, cards[1])

			return HandTypeThreeOfAKind, threeOfAKind, nil
		case len(pairs[0]) == 4:
			fourOfAKind := []Card{pairs[0][0], pairs[0][1], pairs[0][2], pairs[0][3]}
			for _, card := range pairs[0] {
				cards = removeCards(cards, card)
			}
			sortByRankDesc(cards)
			fourOfAKind = append(fourOfAKind, cards[0])
			return HandTypeFourOfAKind, fourOfAKind, nil
		case len(pairs) == 1:
			pair := []Card{pairs[0][0], pairs[0][1]}
			for _, card := range pairs[0] {
				cards = removeCards(cards, card)
			}
			sortByRankDesc(cards)
			pair = append(pair, cards[0])
			pair = append(pair, cards[1])
			pair = append(pair, cards[2])

			return HandTypePair, pair, nil
		case len(pairs) >= 2:
			twoPair := []Card{pairs[0][0], pairs[0][1], pairs[1][0], pairs[1][1]}
			for _, card := range pairs[0] {
				cards = removeCards(cards, card)
			}
			for _, card := range pairs[1] {
				cards = removeCards(cards, card)
			}
			sortByRankDesc(cards)
			twoPair = append(twoPair, cards[0])

			return HandTypeTwoPair, twoPair, nil
		default:
			return HandTypeUnknown, nil, fmt.Errorf("invalid pairs (pairs: %v)", pairs)
		}
	}

	sortByRankDesc(cards)
	highCard := []Card{cards[0], cards[1], cards[2], cards[3], cards[4]}
	return HandTypeHighCard, highCard, nil
}

// CompareHands returns the winner(s) of the game.
func CompareHands(players []Player, board []Card) ([]Player, error) {
	type maxPlayer struct {
		player    Player
		hand      []Card
		score     HandType
		tiePlayer []Player
	}

	var mp maxPlayer

	for i, player := range players {
		score, hand, err := player.Evaluate(board)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate %s's hand: %w", player.Name, err)
		}

		if mp.hand == nil {
			mp = maxPlayer{
				player:    player,
				hand:      hand,
				score:     score,
				tiePlayer: nil,
			}
			continue
		}

		switch {
		case score > mp.score:
			mp = maxPlayer{
				player:    player,
				hand:      hand,
				score:     score,
				tiePlayer: nil,
			}
		case score < mp.score:
			continue
		default:
			winner, err := compareTieHands(player.Hand, mp.player.Hand, board, score)
			if err != nil {
				return nil, fmt.Errorf("failed to compare tie hands: %w", err)
			}

			switch winner {
			case WinnerPlayer1:
				mp = maxPlayer{
					player:    player,
					hand:      hand,
					score:     score,
					tiePlayer: nil,
				}
			case WinnerTie:
				mp.tiePlayer = append(mp.tiePlayer, players[i])
			case WinnerPlayer2:
			default:
				return nil, fmt.Errorf("unknown winner: %d", winner)
			}
		}
	}

	if len(mp.tiePlayer) == 0 {
		return []Player{mp.player}, nil
	}
	result := make([]Player, 0, len(mp.tiePlayer)+1)
	result = append(result, mp.tiePlayer...)
	result = append(result, mp.player)

	return result, nil
}

func compareTieHands(hand1, hand2 []Card, board []Card, score HandType) (Winner, error) {
	switch score {
	case HandTypeStraightFlush:
		winner, err := breakTieByStraightFlush(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByStraightFlush(): %w", err)
		}
		return winner, nil
	case HandTypeFourOfAKind:
		winner, err := breakTieByFourOfAKind(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByFourOfAKind(): %w", err)
		}
		return winner, nil
	case HandTypeFullHouse:
		winner, err := breakTieByFullHouse(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByFullHouse(): %w", err)
		}
		return winner, nil
	case HandTypeFlush:
		winner, err := breakTieByFlush(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByFlush(): %w", err)
		}
		return winner, nil
	case HandTypeStraight:
		winner, err := breakTieByStraight(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByStraight(): %w", err)
		}
		return winner, nil
	case HandTypeThreeOfAKind:
		winner, err := breakTieByThreeOfAKind(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByThreeOfAKind(): %w", err)
		}
		return winner, nil
	case HandTypeTwoPair:
		winner, err := breakTieByTwoPair(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByTwoPair(): %w", err)
		}
		return winner, nil
	case HandTypePair:
		winner, err := breakTieByPair(hand1, hand2, board)
		if err != nil {
			return WinnerUnknown, fmt.Errorf("breakTieByPair(): %w", err)
		}
		return winner, nil
	default:
		return breakTieByHighCard(hand1, hand2), nil
	}
}

// AllCombinations returns all combinations of k elements from the given slice.
func AllCombinations(cards []Card, k int) [][]Card {
	if k == 0 {
		return [][]Card{{}}
	}
	if len(cards) < k {
		return [][]Card{}
	}
	var ret [][]Card
	for i, c := range cards {
		for _, combo := range AllCombinations(cards[i+1:], k-1) {
			ret = append(ret, append([]Card{c}, combo...))
		}
	}
	return ret
}

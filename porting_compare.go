package poker

import (
	"fmt"
)

// CompareVSMadeHand compares the hand of a player against all possible boards.
func CompareVSMadeHand(p1 Player) error {
	deck := NewDeck()
	for _, c := range p1.Hand {
		deck.removeCard(c)
	}

	for _, board := range AllCombinations(deck.Cards, 5) {
		madehand := NewBestMadeHand(append(p1.Hand, board...))

		score, hand, err := p1.Evaluate(board)
		if err != nil {
			return err
		}

		if madehand.Type() != score {
			fmt.Println("madehand:", madehand.Type(), "score:", score, "hand:", hand, "board:", board, "p1:", p1.Hand)
		}
	}

	return nil
}

// EvaluateEquityByMadeHand returns the equity of each player in the game.
func EvaluateEquityByMadeHand(players []Player) ([]float64, error) {
	return EvaluateEquityByMadeHandWithCommunity(players, nil)
}

func EvaluateEquityByMadeHandWithCommunity(players []Player, community []Card) ([]float64, error) {
	deck := NewDeck()

	for _, p := range players {
		for _, c := range p.Hand {
			deck.removeCard(c)
		}
	}

	for _, c := range community {
		deck.removeCard(c)
	}

	wins := make([]int, len(players))

	var ties int
	var total int
	for _, board := range AllCombinations(deck.Cards, 5-len(community)) {
		winners, err := CompareHandsByMadeHand(players, append(community, board...))
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

// CompareHandsByMadeHand returns the winner(s) of the game.
func CompareHandsByMadeHand(players []Player, board []Card) ([]Player, error) {
	type maxPlayer struct {
		player    Player
		hand      []Card
		score     HandType
		tiePlayer []Player
	}

	var mp maxPlayer

	for i, player := range players {
		md := NewBestMadeHand(append(player.Hand, board...))

		if mp.score == HandTypeUnknown {
			mp = maxPlayer{
				player: Player{
					Name:  player.Name,
					Hand:  player.Hand,
					Score: md.Type(),
				},
				hand:      nil,
				score:     md.Type(),
				tiePlayer: nil,
			}
			continue
		}

		switch {
		case md.Type() > mp.score:
			mp = maxPlayer{
				player: Player{
					Name:  player.Name,
					Hand:  player.Hand,
					Score: md.Type(),
				},
				hand:      nil,
				score:     md.Type(),
				tiePlayer: nil,
			}
		case md.Type() < mp.score:
			continue
		default:
			maxMD := NewBestMadeHand(append(mp.player.Hand, board...))

			switch {
			case md.Power() > maxMD.Power():
				mp = maxPlayer{
					player: Player{
						Name:  player.Name,
						Hand:  player.Hand,
						Score: md.Type(),
					},
					hand:      nil,
					score:     md.Type(),
					tiePlayer: nil,
				}
			case md.Power() == maxMD.Power():
				mp.tiePlayer = append(mp.tiePlayer, players[i])
			case md.Power() < maxMD.Power():
				// do nothing
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

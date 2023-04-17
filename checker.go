package poker

import "sort"

func IsStraightFlush(cards []Card) []Card {
	flush, suit := isFlush(cards)
	if flush == nil {
		return nil
	}

	suitCards := getSuitCards(cards, *suit)
	straightFlush := IsStraight(suitCards)
	if straightFlush == nil {
		return nil
	}
	return *straightFlush
}

func isFlush(cards []Card) ([]Card, *Suit) {
	suits := make(map[Suit]int)
	for _, card := range cards {
		suits[card.Suit]++
	}
	for suit, count := range suits {
		if count >= 5 {
			f := getFlush(cards, suit)
			return f, &suit
		}
	}
	return nil, nil
}

func getFlush(cards []Card, suit Suit) []Card {
	suitCards := getSuitCards(cards, suit)

	return suitCards[:5]
}

func getSuitCards(cards []Card, suit Suit) []Card {
	suitCards := make([]Card, 0)
	for _, card := range cards {
		if card.Suit == suit {
			suitCards = append(suitCards, card)
		}
	}
	sort.Slice(suitCards, func(i, j int) bool {
		return suitCards[i].Rank > suitCards[j].Rank
	})
	return suitCards
}

func IsStraight(cards []Card) *[]Card {
	cards = uniqueCards(cards)

	cards = sortByRank(cards)
	// Straight is staring three in seven cards
	for i := 0; i < 3; i++ {
		count := 0
		for j := i; j < len(cards)-1; j++ {
			if (j + 1) >= len(cards) {
				break
			}

			if cards[j].Rank == cards[j+1].Rank-1 {
				count++
			} else {
				break
			}
		}

		// if count is 4, it is straight
		if count == 4 {
			return &[]Card{cards[i+4], cards[i+3], cards[i+2], cards[i+1], cards[i]}
		}
	}

	// A to 5 is straight
	cards = sortByRank(cards)
	if cards[0].Rank == RankDeuce {
		count := 0
		for i := 0; i < len(cards)-1; i++ {
			if cards[i].Rank == cards[i+1].Rank-1 {
				count++
			} else {
				break
			}
		}

		if count == 3 && cards[len(cards)-1].Rank == RankAce {
			return &[]Card{cards[3], cards[2], cards[1], cards[0], cards[len(cards)-1]}
		}
	}

	return nil
}

func GetPairs(cards []Card) [][]Card {
	ranks := make(map[Rank][]int)
	for i, card := range cards {
		ranks[card.Rank] = append(ranks[card.Rank], i)
	}
	pairs := make([][]Card, 0)
	for rank, indexes := range ranks {
		if len(indexes) >= 2 {
			p := make([]Card, 0)
			for _, card := range cards {
				if card.Rank == rank {
					p = append(p, card)
				}
			}

			pairs = append(pairs, p)
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0].Rank > pairs[j][0].Rank
	})
	sort.Slice(pairs, func(i, j int) bool {
		return len(pairs[i]) > len(pairs[j])
	})

	return pairs
}

func isFullHouse(pairs [][]Card) []Card {
	if len(pairs) <= 1 {
		return nil
	}

	sort.Slice(pairs, func(i, j int) bool {
		if len(pairs[i]) > len(pairs[j]) {
			return true
		}

		if len(pairs[i]) == len(pairs[j]) {
			return pairs[i][0].Rank > pairs[j][0].Rank
		}
		return len(pairs[i]) > len(pairs[j])
	})

	if len(pairs[0]) == 2 {
		pairs[0], pairs[1] = pairs[1], pairs[0]
	}
	if len(pairs[0]) != 3 || len(pairs[1]) <= 1 {
		return nil
	}
	return []Card{pairs[0][0], pairs[0][1], pairs[0][2], pairs[1][0], pairs[1][1]}
}

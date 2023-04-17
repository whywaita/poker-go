package poker

import "fmt"

func breakTieByHighCard(hand1, hand2 []Card) Winner {
	hand1 = sortByRank(hand1)
	hand2 = sortByRank(hand2)

	for i := 0; i < len(hand1); i++ {
		if hand1[i].Rank > hand2[i].Rank {
			return WinnerPlayer1
		} else if hand1[i].Rank < hand2[i].Rank {
			return WinnerPlayer2
		}
	}

	return WinnerTie
}

func breakTieByHighCardWithoutPair(pairs1, pairs2 [][]Card, hand1, hand2 []Card) Winner {
	// Evaluate kicker cards without pair cards
	// Remove pair cards from hand
	switch {
	case len(pairs1) == 1:
		for _, c := range pairs1[0] {
			hand1 = removeCards(hand1, c)
		}
		for _, c := range pairs2[0] {
			hand1 = removeCards(hand1, c)
		}
	case len(pairs1) == 2:
		for i := 0; i < 1; i++ {
			for _, c := range pairs1[i] {
				hand1 = removeCards(hand1, c)
			}
		}
		for i := 0; i < 1; i++ {
			for _, c := range pairs2[i] {
				hand2 = removeCards(hand2, c)
			}
		}
	}

	return breakTieByHighCard(hand1, hand2)
}

func breakTieByPair(hand1, hand2 []Card, board []Card) (Winner, error) {
	h1, h2 := append(hand1, board...), append(hand2, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1) != 1 || len(pairs2) != 1 {
		return WinnerUnknown, fmt.Errorf("input is not one pair (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return WinnerPlayer1, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return WinnerPlayer2, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2), nil
	}
}

func breakTieByTwoPair(hand1, hand2 []Card, board []Card) (Winner, error) {
	h1, h2 := append(hand1, board...), append(hand2, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1) <= 1 || len(pairs2) <= 1 {
		return WinnerUnknown, fmt.Errorf("input is not two pair (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return WinnerPlayer1, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return WinnerPlayer2, nil
	} else if pairs1[1][0].Rank > pairs2[1][0].Rank {
		return WinnerPlayer1, nil
	} else if pairs1[1][0].Rank < pairs2[1][0].Rank {
		return WinnerPlayer2, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2), nil
	}
}

func breakTieByThreeOfAKind(hand1, hand2 []Card, board []Card) (Winner, error) {
	h1, h2 := append(hand1, board...), append(hand2, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1[0]) != 3 || len(pairs2[0]) != 3 {
		return WinnerUnknown, fmt.Errorf("input is not three of a kind (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return WinnerPlayer1, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return WinnerPlayer2, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2), nil
	}
}

func breakTieByStraight(hand1, hand2 []Card, board []Card) (Winner, error) {
	h1, h2 := append(hand1, board...), append(hand2, board...)
	s1 := IsStraight(h1)
	s2 := IsStraight(h2)

	if s1 == nil {
		return WinnerUnknown, fmt.Errorf("input is not straight (s1: %v, h1: %v board: %v)", s1, h1, board)
	}
	if s2 == nil {
		return WinnerUnknown, fmt.Errorf("input is not straight (s2: %v, h2: %v board: %v)", s2, h2, board)
	}

	straight1 := sortByRank(*s1)
	straight2 := sortByRank(*s2)

	if straight1[0].Rank > straight2[0].Rank {
		return WinnerPlayer1, nil
	} else if straight1[0].Rank < straight2[0].Rank {
		return WinnerPlayer2, nil
	} else {
		return WinnerTie, nil
	}
}

func breakTieByFlush(hand1, hand2 []Card, board []Card) (Winner, error) {
	flush1, _ := isFlush(append(hand1, board...))
	flush2, _ := isFlush(append(hand2, board...))

	if flush1 == nil || flush2 == nil {
		return WinnerUnknown, fmt.Errorf("input is not flush (p1: %v, p2: %v, board: %v)", flush1, flush2, board)
	}

	for i := 0; i < len(flush1); i++ {
		if flush1[i].Rank > flush2[i].Rank {
			return WinnerPlayer1, nil
		} else if flush1[i].Rank < flush2[i].Rank {
			return WinnerPlayer2, nil
		}
	}

	return WinnerTie, nil
}

func breakTieByFullHouse(hand1, hand2 []Card, board []Card) (Winner, error) {
	h1, h2 := append(hand1, board...), append(hand2, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	fullHouse1 := isFullHouse(pairs1)
	fullHouse2 := isFullHouse(pairs2)
	if fullHouse1 == nil || fullHouse2 == nil {
		return WinnerUnknown, fmt.Errorf("input is not full house (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}

	if fullHouse1[0].Rank > fullHouse2[0].Rank {
		return WinnerPlayer1, nil
	} else if fullHouse1[0].Rank < fullHouse2[0].Rank {
		return WinnerPlayer2, nil
	} else if fullHouse1[3].Rank > fullHouse2[3].Rank {
		return WinnerPlayer1, nil
	} else if fullHouse1[3].Rank < fullHouse2[3].Rank {
		return WinnerPlayer2, nil
	} else {
		return WinnerTie, nil
	}
}

func breakTieByFourOfAKind(hand1, hand2 []Card, board []Card) (Winner, error) {
	h1, h2 := append(hand1, board...), append(hand2, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1[0]) != 4 || len(pairs2[0]) != 4 {
		return WinnerUnknown, fmt.Errorf("input is not four of a kind (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return WinnerPlayer1, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return WinnerPlayer2, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2), nil
	}
}

func breakTieByStraightFlush(hand1, hand2 []Card, board []Card) (Winner, error) {
	sf1 := IsStraightFlush(append(hand1, board...))
	sf2 := IsStraightFlush(append(hand2, board...))

	if sf1 == nil || sf2 == nil {
		return WinnerUnknown, fmt.Errorf("input is not straight flush (p1: %v, p2: %v, board: %v)", sf1, sf2, board)
	}

	straightFlush1 := sortByRank(sf1)
	straightFlush2 := sortByRank(sf2)

	if straightFlush1[0].Rank > straightFlush2[0].Rank {
		return WinnerPlayer1, nil
	} else if straightFlush1[0].Rank < straightFlush2[0].Rank {
		return WinnerPlayer2, nil
	} else {
		return WinnerTie, nil
	}
}

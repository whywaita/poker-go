package poker

import "sort"

func sortByRankDesc(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank > cards[j].Rank
	})
	return cards
}

func sortByRank(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})
	return cards
}

func uniqueCards(cards []Card) []Card {
	unique := make([]Card, 0, len(cards))
	seen := make(map[Rank]bool)
	for _, c := range cards {
		if !seen[c.Rank] {
			unique = append(unique, c)
			seen[c.Rank] = true
		}
	}
	return unique
}

func indexOf(element Player, data []Player) int {
	for k, v := range data {
		if element.Name == v.Name {
			return k
		}
	}
	return -1 //not found.
}

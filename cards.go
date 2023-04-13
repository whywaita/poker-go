package poker

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Rank int
	Suit Suit
}

type Deck struct {
	Cards []Card
}

type Suit int

const (
	Hearts Suit = iota
	Clubs
	Diamonds
	Spades
)

func (s Suit) String() string {
	switch s {
	case Hearts:
		return "hearts"
	case Clubs:
		return "clubs"
	case Diamonds:
		return "diamonds"
	case Spades:
		return "spades"
	default:
		return "unknown"
	}
}

//
//type Hand struct {
//	Cards []Card
//}

func NewDeck() *Deck {
	d := &Deck{Cards: make([]Card, 52)}
	for i := 0; i < 13; i++ {
		d.Cards[i] = Card{i + 2, Hearts}
		d.Cards[i+13] = Card{i + 2, Clubs}
		d.Cards[i+26] = Card{i + 2, Diamonds}
		d.Cards[i+39] = Card{i + 2, Spades}
	}
	return d
}

func (d *Deck) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) DrawCard() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) DrawCards(n int) []Card {
	cards := make([]Card, n)
	for i := 0; i < n; i++ {
		cards[i] = d.DrawCard()
	}
	return cards
}

//func (d *Deck) RemoveCards(cardString string) error {
//	for _, c := range cardString {
//		card := Card{Rank: int(c - '0'), Suit: string(cardString[len(cardString)-1])}
//		if !d.removeCard(card) {
//			return fmt.Errorf("%v is not found", card)
//		}
//	}
//	return nil
//}

func (d *Deck) removeCard(card Card) bool {
	for i, c := range d.Cards {
		if c.Rank == card.Rank && c.Suit == card.Suit {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
			return true
		}
	}
	return false
}

func removeCards(cards []Card, c Card) []Card {
	for i, card := range cards {
		if card.Rank == c.Rank && card.Suit == c.Suit {
			return append(cards[:i], cards[i+1:]...)
		}
	}
	return cards
}

//func NewHand(cards []Card) *Hand {
//	return &Hand{Cards: cards}
//}
//
//func (h *Hand) AddCard(card Card) {
//	h.Cards = append(h.Cards, card)
//}
//
//func (h *Hand) RemoveCard(card Card) bool {
//	for i, c := range h.Cards {
//		if c.Rank == card.Rank && c.Suit == card.Suit {
//			h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
//			return true
//		}
//	}
//	return false
//}
//
//func (h *Hand) String() string {
//	str := ""
//	for _, c := range h.Cards {
//		str += fmt.Sprintf("%d%s ", c.Rank, c.Suit)
//	}
//	return str
//}

type Player struct {
	Name  string
	Hand  []Card
	Score int
}

func NewPlayer(name string, hand []Card) *Player {
	return &Player{
		Name: name,
		Hand: hand,
	}
}

func (p *Player) Evaluate(board []Card) (Hand, []Card, error) {
	cards := append(p.Hand, board...)
	return Evaluate(cards)
}

func Evaluate(cards []Card) (Hand, []Card, error) {
	if len(cards) != 7 {
		return 0, nil, fmt.Errorf("invalid number of cards")
	}
	return evaluate(cards)
}

type Hand int

const (
	HandRoyalFlush    Hand = 10
	HandStraightFlush Hand = 9
	HandFourOfAKind   Hand = 8
	HandFullHouse     Hand = 7
	HandFlush         Hand = 6
	HandStraight      Hand = 5
	HandThreeOfAKind  Hand = 4
	HandTwoPair       Hand = 3
	HandPair          Hand = 2
	HandHighCard      Hand = 1
	HandUnknown       Hand = 0
)

func (h Hand) String() string {
	switch h {
	case HandRoyalFlush:
		return "Royal Flush"
	case HandStraightFlush:
		return "Straight Flush"
	case HandFourOfAKind:
		return "Four of a Kind"
	case HandFullHouse:
		return "Full House"
	case HandFlush:
		return "Flush"
	case HandStraight:
		return "Straight"
	case HandThreeOfAKind:
		return "Three of a Kind"
	case HandTwoPair:
		return "Two Pair"
	case HandPair:
		return "Pair"
	case HandHighCard:
		return "High Card"
	}
	return "unknown"
}

func evaluate(cards []Card) (Hand, []Card, error) {
	cards = sortByRankDesc(cards)

	flush, _ := isFlush(cards)
	straight := IsStraight(cards)
	straightFlush := IsStraightFlush(cards)
	pairs := GetPairs(cards)
	fullHouse := isFullHouse(pairs)

	if straightFlush != nil {
		sortByRank(straightFlush)
		if straightFlush[4].Rank == 14 {
			return HandRoyalFlush, straightFlush, nil
		}
		return HandStraightFlush, straightFlush, nil
	} else if fullHouse != nil {
		return HandFullHouse, fullHouse, nil
	} else if flush != nil {
		return HandFlush, flush, nil
	} else if straight != nil {
		return HandStraight, *straight, nil
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

			return HandThreeOfAKind, threeOfAKind, nil
		case len(pairs[0]) == 4:
			fourOfAKind := []Card{pairs[0][0], pairs[0][1], pairs[0][2], pairs[0][3]}
			for _, card := range pairs[0] {
				cards = removeCards(cards, card)
			}
			sortByRankDesc(cards)
			fourOfAKind = append(fourOfAKind, cards[0])
			return HandFourOfAKind, fourOfAKind, nil
		case len(pairs) == 1:
			pair := []Card{pairs[0][0], pairs[0][1]}
			for _, card := range pairs[0] {
				cards = removeCards(cards, card)
			}
			sortByRankDesc(cards)
			pair = append(pair, cards[0])
			pair = append(pair, cards[1])
			pair = append(pair, cards[2])

			return HandPair, pair, nil
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

			return HandTwoPair, twoPair, nil
		default:
			return HandUnknown, nil, fmt.Errorf("invalid pairs (pairs: %v)", pairs)
		}
	}

	sortByRankDesc(cards)
	highCard := []Card{cards[0], cards[1], cards[2], cards[3], cards[4]}
	return HandHighCard, highCard, nil
}

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
	cards = sortByRank(cards)
	if cards[0].Rank == 14 && cards[1].Rank == 5 && cards[2].Rank == 4 && cards[3].Rank == 3 && cards[4].Rank == 2 {
		return &[]Card{cards[0], cards[1], cards[2], cards[3], cards[4]}
	}

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
			return &[]Card{cards[i], cards[i+1], cards[i+2], cards[i+3], cards[i+4]}
		}
	}

	// A to 5 is straight
	cards = sortByRankDesc(cards)
	if cards[0].Rank == 14 {
		cards = sortByRank(cards)
		for i := 0; i < len(cards)-1; i++ {
			if cards[i].Rank == 2 {
				count := 0
				for j := i; j < i+5; j++ {
					if cards[j].Rank == cards[j+1].Rank-1 {
						count++
					} else {
						break
					}
				}
				if count == 3 {
					return &[]Card{cards[i], cards[i+1], cards[i+2], cards[i+3], cards[len(cards)-1]}
				}
			}
		}
	}

	return nil
}

func GetPairs(cards []Card) [][]Card {
	ranks := make(map[int][]int)
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

func CompareHands(hand1 []Card, hand2 []Card, board []Card) (string, error) {
	player1 := NewPlayer("player1", hand1)
	player2 := NewPlayer("player2", hand2)
	score1, _, err := player1.Evaluate(board)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate %s's hand: %w", player1.Name, err)
	}
	score2, _, err := player2.Evaluate(board)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate %s's hand: %w", player2.Name, err)
	}

	if score1 == HandThreeOfAKind && score2 == HandTwoPair {
		fmt.Printf("score1 %s | score2 %s | player1 %v | player2 %v | board %v\n", score1, score2, player1.Hand, player2.Hand, board)
	}

	if score1 > score2 {
		return player1.Name, nil
	} else if score1 < score2 {
		return player2.Name, nil
	} else {
		// If there is a tie, break it by hand rank
		switch score1 {
		case HandStraightFlush:
			winner, err := breakTieByStraightFlush(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByStraightFlush(): %w", err)
			}
			return winner, nil
		case HandFourOfAKind:
			winner, err := breakTieByFourOfAKind(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByFourOfAKind(): %w", err)
			}
			return winner, nil
		case HandFullHouse:
			winner, err := breakTieByFullHouse(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByFullHouse(): %w", err)
			}
			return winner, nil
		case HandFlush:
			winner, err := breakTieByFlush(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByFlush(): %w", err)
			}
			return winner, nil
		case HandStraight:
			winner, err := breakTieByStraight(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByStraight(): %w", err)
			}
			return winner, nil
		case HandThreeOfAKind:
			winner, err := breakTieByThreeOfAKind(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByThreeOfAKind(): %w", err)
			}
			return winner, nil
		case HandTwoPair:
			winner, err := breakTieByTwoPair(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByTwoPair(): %w", err)
			}
			return winner, nil
		case HandPair:
			winner, err := breakTieByPair(*player1, *player2, board)
			if err != nil {
				return "", fmt.Errorf("breakTieByPair(): %w", err)
			}
			return winner, nil
		default:
			return breakTieByHighCard(*player1, *player2), nil
		}
	}
}

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

func breakTieByHighCard(p1, p2 Player) string {
	for i := 0; i < len(p1.Hand); i++ {
		if p1.Hand[i].Rank > p2.Hand[i].Rank {
			return p1.Name
		} else if p1.Hand[i].Rank < p2.Hand[i].Rank {
			return p2.Name
		}
	}

	return "tie"
}

func breakTieByHighCardWithoutPair(pairs1, pairs2 [][]Card, hand1, hand2 []Card, name1, name2 string) (string, error) {
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

	// Sort by rank
	hand1 = sortByRank(hand1)
	hand2 = sortByRank(hand2)

	p1 := NewPlayer(name1, hand1)
	p2 := NewPlayer(name2, hand2)

	return breakTieByHighCard(*p1, *p2), nil
}

func breakTieByPair(p1, p2 Player, board []Card) (string, error) {
	h1, h2 := append(p1.Hand, board...), append(p2.Hand, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1) != 1 || len(pairs2) != 1 {
		return "", fmt.Errorf("input is not one pair (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return p1.Name, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return p2.Name, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2, p1.Name, p2.Name)
	}
}

func breakTieByTwoPair(p1, p2 Player, board []Card) (string, error) {
	h1, h2 := append(p1.Hand, board...), append(p2.Hand, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1) <= 1 || len(pairs2) <= 1 {
		return "", fmt.Errorf("input is not two pair (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return p1.Name, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return p2.Name, nil
	} else if pairs1[1][0].Rank > pairs2[1][0].Rank {
		return p1.Name, nil
	} else if pairs1[1][0].Rank < pairs2[1][0].Rank {
		return p2.Name, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2, p1.Name, p2.Name)
	}
}

func breakTieByThreeOfAKind(p1, p2 Player, board []Card) (string, error) {
	h1, h2 := append(p1.Hand, board...), append(p2.Hand, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1[0]) != 3 || len(pairs2[0]) != 3 {
		return "", fmt.Errorf("input is not three of a kind (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return p1.Name, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return p2.Name, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2, p1.Name, p2.Name)
	}
}

func breakTieByStraight(p1, p2 Player, board []Card) (string, error) {
	h1, h2 := append(p1.Hand, board...), append(p2.Hand, board...)
	s1 := IsStraight(h1)
	s2 := IsStraight(h2)

	if s1 == nil {
		return "", fmt.Errorf("input is not straight (s1: %v, h1: %v board: %v)", s1, h1, board)
	}
	if s2 == nil {
		return "", fmt.Errorf("input is not straight (s2: %v, h2: %v board: %v)", s2, h2, board)
	}

	straight1 := sortByRank(*s1)
	straight2 := sortByRank(*s2)

	if straight1[0].Rank > straight2[0].Rank {
		return p1.Name, nil
	} else if straight1[0].Rank < straight2[0].Rank {
		return p2.Name, nil
	} else {
		return "tie", nil
	}
}

func breakTieByFlush(p1, p2 Player, board []Card) (string, error) {
	flush1, _ := isFlush(append(p1.Hand, board...))
	flush2, _ := isFlush(append(p2.Hand, board...))

	if flush1 == nil || flush2 == nil {
		return "", fmt.Errorf("input is not flush (p1: %v, p2: %v, board: %v)", flush1, flush2, board)
	}

	for i := 0; i < len(flush1); i++ {
		if flush1[i].Rank > flush2[i].Rank {
			return p1.Name, nil
		} else if flush1[i].Rank < flush2[i].Rank {
			return p2.Name, nil
		}
	}

	return "tie", nil
}

func breakTieByFullHouse(p1, p2 Player, board []Card) (string, error) {
	h1, h2 := append(p1.Hand, board...), append(p2.Hand, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	fullHouse1 := isFullHouse(pairs1)
	fullHouse2 := isFullHouse(pairs2)
	if fullHouse1 == nil || fullHouse2 == nil {
		return "", fmt.Errorf("input is not full house (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}

	if fullHouse1[0].Rank > fullHouse2[0].Rank {
		return p1.Name, nil
	} else if fullHouse1[0].Rank < fullHouse2[0].Rank {
		return p2.Name, nil
	} else if fullHouse1[3].Rank > fullHouse2[3].Rank {
		return p1.Name, nil
	} else if fullHouse1[3].Rank < fullHouse2[3].Rank {
		return p2.Name, nil
	} else {
		return "tie", nil
	}
}

func breakTieByFourOfAKind(p1, p2 Player, board []Card) (string, error) {
	h1, h2 := append(p1.Hand, board...), append(p2.Hand, board...)
	pairs1 := GetPairs(h1)
	pairs2 := GetPairs(h2)
	if len(pairs1[0]) != 4 || len(pairs2[0]) != 4 {
		return "", fmt.Errorf("input is not four of a kind (p1: %v, p2: %v, board: %v)", pairs1, pairs2, board)
	}
	if pairs1[0][0].Rank > pairs2[0][0].Rank {
		return p1.Name, nil
	} else if pairs1[0][0].Rank < pairs2[0][0].Rank {
		return p2.Name, nil
	} else {
		return breakTieByHighCardWithoutPair(pairs1, pairs2, h1, h2, p1.Name, p2.Name)
	}
}

func breakTieByStraightFlush(p1, p2 Player, board []Card) (string, error) {
	sf1 := IsStraightFlush(append(p1.Hand, board...))
	sf2 := IsStraightFlush(append(p2.Hand, board...))

	if sf1 == nil || sf2 == nil {
		return "", fmt.Errorf("input is not straight flush (p1: %v, p2: %v, board: %v)", sf1, sf2, board)
	}

	straightFlush1 := sortByRank(sf1)
	straightFlush2 := sortByRank(sf2)

	if straightFlush1[0].Rank > straightFlush2[0].Rank {
		return p1.Name, nil
	} else if straightFlush1[0].Rank < straightFlush2[0].Rank {
		return p2.Name, nil
	} else {
		return "tie", nil
	}
}

func EvaluateEquity(hand1 []Card, hand2 []Card) (float64, float64, error) {
	deck := NewDeck()
	for _, c := range hand1 {
		deck.removeCard(c)
	}
	for _, c := range hand2 {
		deck.removeCard(c)
	}
	player1 := NewPlayer("player1", hand1)
	player2 := NewPlayer("player2", hand2)
	var player1Wins int
	var player2Wins int
	var ties int
	for _, board := range AllCombinations(deck.Cards, 5) {
		winner, err := CompareHands(player1.Hand, player2.Hand, board)
		if err != nil {
			return 0, 0, err
		}
		//fmt.Println(winner)
		switch winner {
		case "player1":
			player1Wins++
		case "player2":
			player2Wins++
		case "tie":
			ties++
		}
	}
	fmt.Println("player1 wins:", player1Wins)
	fmt.Println("player2 wins:", player2Wins)
	fmt.Println("ties:", ties)
	total := player1Wins + player2Wins + ties
	fmt.Println("total:", total)
	return float64(player1Wins+(ties/2)) / float64(total), float64(player2Wins+(ties/2)) / float64(total), nil
}

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

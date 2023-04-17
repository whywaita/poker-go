package poker

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Rank Rank
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

type Rank int

const (
	RankUnknown Rank = iota
	RankDeuce        = iota
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
)

func (r Rank) String() string {
	switch r {
	case RankDeuce:
		return "2"
	case RankThree:
		return "3"
	case RankFour:
		return "4"
	case RankFive:
		return "5"
	case RankSix:
		return "6"
	case RankSeven:
		return "7"
	case RankEight:
		return "8"
	case RankNine:
		return "9"
	case RankTen:
		return "T"
	case RankJack:
		return "J"
	case RankQueen:
		return "Q"
	case RankKing:
		return "K"
	case RankAce:
		return "A"
	default:
		return "Unknown"
	}
}

func UnmarshalRank(r any) Rank {
	switch r := r.(type) {
	case string:
		return UnmarshalRankString(r)
	case int:
		return UnmarshalRankInt(r)
	default:
		return RankUnknown
	}
}

func UnmarshalRankString(r string) Rank {
	switch r {
	case "2":
		return RankDeuce
	case "3":
		return RankThree
	case "4":
		return RankFour
	case "5":
		return RankFive
	case "6":
		return RankSix
	case "7":
		return RankSeven
	case "8":
		return RankEight
	case "9":
		return RankNine
	case "T":
		return RankTen
	case "J":
		return RankJack
	case "Q":
		return RankQueen
	case "K":
		return RankKing
	case "A":
		return RankAce
	default:
		return RankUnknown
	}
}

func UnmarshalRankInt(r int) Rank {
	switch r {
	case 2:
		return RankDeuce
	case 3:
		return RankThree
	case 4:
		return RankFour
	case 5:
		return RankFive
	case 6:
		return RankSix
	case 7:
		return RankSeven
	case 8:
		return RankEight
	case 9:
		return RankNine
	case 10:
		return RankTen
	case 11:
		return RankJack
	case 12:
		return RankQueen
	case 13:
		return RankKing
	case 14:
		return RankAce
	default:
		return RankUnknown
	}
}

func NewDeck() *Deck {
	d := &Deck{Cards: make([]Card, 52)}
	for i := 0; i < 13; i++ {
		d.Cards[i] = Card{UnmarshalRank(i + 2), Hearts}
		d.Cards[i+13] = Card{UnmarshalRank(i + 2), Clubs}
		d.Cards[i+26] = Card{UnmarshalRank(i + 2), Diamonds}
		d.Cards[i+39] = Card{UnmarshalRank(i + 2), Spades}
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

func (p *Player) Evaluate(board []Card) (HandType, []Card, error) {
	cards := append(p.Hand, board...)
	return Evaluate(cards)
}

func Evaluate(cards []Card) (HandType, []Card, error) {
	if len(cards) != 7 {
		return 0, nil, fmt.Errorf("invalid number of cards")
	}
	return evaluate(cards)
}

type HandType int

const (
	HandTypeRoyalFlush    HandType = 10
	HandTypeStraightFlush HandType = 9
	HandTypeFourOfAKind   HandType = 8
	HandTypeFullHouse     HandType = 7
	HandTypeFlush         HandType = 6
	HandTypeStraight      HandType = 5
	HandTypeThreeOfAKind  HandType = 4
	HandTypeTwoPair       HandType = 3
	HandTypePair          HandType = 2
	HandTypeHighCard      HandType = 1
	HandTypeUnknown       HandType = 0
)

func (h HandType) String() string {
	switch h {
	case HandTypeRoyalFlush:
		return "Royal Flush"
	case HandTypeStraightFlush:
		return "Straight Flush"
	case HandTypeFourOfAKind:
		return "Four of a Kind"
	case HandTypeFullHouse:
		return "Full House"
	case HandTypeFlush:
		return "Flush"
	case HandTypeStraight:
		return "Straight"
	case HandTypeThreeOfAKind:
		return "Three of a Kind"
	case HandTypeTwoPair:
		return "Two Pair"
	case HandTypePair:
		return "Pair"
	case HandTypeHighCard:
		return "High Card"
	}
	return "unknown"
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

type Winner int

const (
	WinnerUnknown Winner = iota
	WinnerPlayer1
	WinnerPlayer2
	WinnerTie
)

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

func indexOf(element Player, data []Player) int {
	for k, v := range data {
		if element.Name == v.Name {
			return k
		}
	}
	return -1 //not found.
}

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

func CompareHandsByMadeHand(player1, player2 Player, board []Card) string {
	madehand1 := NewBestMadeHand(append(player1.Hand, board...))
	madehand2 := NewBestMadeHand(append(player2.Hand, board...))

	switch {
	case madehand1.Power() > madehand2.Power():
		return player1.Name
	case madehand1.Power() < madehand2.Power():
		return player2.Name
	default:
		return "tie"
	}
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

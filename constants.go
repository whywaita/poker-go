package poker

import (
	"math/rand"
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

func UnmarshalSuitString(s string) Suit {
	switch s {
	case "hearts", "h":
		return Hearts
	case "clubs", "c":
		return Clubs
	case "diamonds", "d":
		return Diamonds
	case "spades", "s":
		return Spades
	default:
		return -1
	}
}

type Rank int

const (
	RankUnknown Rank = iota
	RankDeuce
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

type Player struct {
	Name  string
	Hand  []Card
	Score HandType
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

type Winner int

const (
	WinnerUnknown Winner = iota
	WinnerPlayer1
	WinnerPlayer2
	WinnerTie
)

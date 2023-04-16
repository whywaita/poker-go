//go:build wasm

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"syscall/js"

	"github.com/whywaita/poker-go"
)

func GenerateHands(this js.Value, args []js.Value) any {
	cards := generateCards()
	hand, handCards, err := poker.Evaluate(cards)
	if err != nil {
		return map[string]any{
			"ok":      false,
			"message": err.Error(),
		}
	}

	notUsedCards := make([]poker.Card, 0)
	for _, card := range cards {
		if !isExist(handCards, card) {
			notUsedCards = append(notUsedCards, card)
		}
	}

	jsCards := convertJS(handCards)
	cardJson, err := json.Marshal(jsCards)
	if err != nil {
		return map[string]any{
			"ok":      false,
			"message": err.Error(),
		}
	}
	jsNotUsedCards := convertJS(notUsedCards)
	notUsedCardsJson, err := json.Marshal(jsNotUsedCards)
	if err != nil {
		return map[string]any{
			"ok":      false,
			"message": err.Error(),
		}
	}

	return map[string]any{
		"ok":           true,
		"handCards":    string(cardJson),
		"notUsedCards": string(notUsedCardsJson),
		"hand":         hand.String(),
	}
}

type jsCard struct {
	Rank string `json:"rank"`
	Suit string `json:"suit"`
}

func convertJS(cards []poker.Card) []jsCard {
	jsCards := make([]jsCard, 0)
	for _, card := range cards {
		jc := jsCard{}
		switch card.Rank {
		case 14:
			jc.Rank = "ace"
		case 13:
			jc.Rank = "king"
		case 12:
			jc.Rank = "queen"
		case 11:
			jc.Rank = "jack"
		default:
			jc.Rank = fmt.Sprintf("%d", card.Rank)
		}
		jc.Suit = card.Suit.String()
		jsCards = append(jsCards, jc)
	}
	return jsCards
}

func generateCards() []poker.Card {
	cards := make([]poker.Card, 0)
	for i := 0; i < 30; i++ {
		rank := rand.Intn(13) + 2
		suit := rand.Intn(4)
		card := poker.Card{Rank: poker.UnmarshalRankInt(rank), Suit: poker.Suit(suit)}
		if !isExist(cards, card) {
			cards = append(cards, card)
			if len(cards) == 7 {
				break
			}
		}
	}
	return cards
}

func isExist(cards []poker.Card, card poker.Card) bool {
	for _, c := range cards {
		if c.Rank == card.Rank && c.Suit == card.Suit {
			return true
		}
	}
	return false
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("poker", js.ValueOf(map[string]any{
		"generateHands": js.FuncOf(GenerateHands),
	}))

	<-c
}

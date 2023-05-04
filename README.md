# poker-go

poker-go is a library for playing poker in Go.

## Usage

- Evaluate a hand: [./examples/eval_hand.go](./examples/eval_hand.go)

```go
package main

import (
	"fmt"
    "log"
	
    "github.com/whywaita/poker-go"
)

func main() {
	h1 := []poker.Card{
		{Rank: poker.RankAce, Suit: poker.Hearts},
		{Rank: poker.RankAce, Suit: poker.Diamonds},
	}
	p1 := poker.NewPlayer("player1", h1)

	board := []poker.Card{
		{Rank: poker.RankAce, Suit: poker.Clubs},
		{Rank: poker.RankAce, Suit: poker.Spades},
		{Rank: poker.RankKing, Suit: poker.Hearts},
		{Rank: poker.RankQueen, Suit: poker.Hearts},
		{Rank: poker.RankJack, Suit: poker.Hearts},
	}

	handtype, cards, err := poker.Evaluate(append(p1.Hand, board...))
	if err != nil {
		log.Fatalf("poker.Evaluate(append(%s. %s...)): %v", p1.Hand, board, err)
	}

	fmt.Printf("handtype: %s, cards: %s\n", handtype, cards)
}
```

```bash
$ go run eval_hand.go
handtype: Four of a Kind, cards: [{A hearts} {A diamonds} {A clubs} {A spades} {K hearts}]
```

- Calculate an equities for a given hands: [./examples/calc_equities.go](./examples/calc_equities)

```go
package main

import (
    "fmt"
    "log"

    "github.com/whywaita/poker-go"
)

func main() {
	h1 := []poker.Card{
		{Rank: poker.RankDeuce, Suit: poker.Hearts},
		{Rank: poker.RankThree, Suit: poker.Diamonds},
	}
	p1 := poker.NewPlayer("player1", h1)

	h2 := []poker.Card{
		{Rank: poker.RankAce, Suit: poker.Hearts},
		{Rank: poker.RankAce, Suit: poker.Diamonds},
	}
	p2 := poker.NewPlayer("player2", h2)

	h3 := []poker.Card{
		{Rank: poker.RankSeven, Suit: poker.Clubs},
		{Rank: poker.RankEight, Suit: poker.Clubs},
	}
	p3 := poker.NewPlayer("player3", h3)

	equities, err := poker.EvaluateEquity([]poker.Player{*p1, *p2, *p3})
	if err != nil {
		log.Fatalf("failed to evaluate equity: %v", err)
	}
	fmt.Printf("%v equity: %f, %v equity: %f, %v equity %f\n", h1, equities[0], h2, equities[1], h3, equities[2])
}
```

```bash
$ go run calc_equities.go
[{2 hearts} {3 diamonds}] equity: 0.085326, [{A hearts} {A diamonds}] equity: 0.663249, [{7 clubs} {8 clubs}] equity 0.251425
```

- Full documents are available at [pkg.go.dev](https://pkg.go.dev/github.com/whywaita/poker-go)

## LICENSE

MIT License

`porting_made_hand.go` and `porting_precalculated_table.go` is porting from [axross/poker](https://github.com/axross/poker).
Only these files are under Apache License 2.0 from the original works.
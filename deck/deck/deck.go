package deck

import (
	"fmt"
	"strings"
)

// Color of a card
type Color string

// Suit of a card
type Suit string

// Suits of cards
const (
	Spades   Suit = "spades"
	Clubs    Suit = "clubs"
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
)

// Colors of cards
const (
	Black Color = "black"
	Red   Color = "red"
)

// Rank of a card
type Rank uint8

// Ranks of cards
const (
	Ace   Rank = 1
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
)

var ranks = map[Rank]string{
	Ace:   "Ace",
	Two:   "Two",
	Three: "Three",
	Four:  "Four",
	Five:  "Five",
	Six:   "Six",
	Seven: "Seven",
	Eight: "Eight",
	Nine:  "Nine",
	Ten:   "Ten",
	Jack:  "Jack",
	Queen: "Queen",
	King:  "King",
}

func (r Rank) String() string {
	rank, ok := ranks[r]
	if !ok {
		return ""
	}
	return rank
}

// Card represents a normal playing card with a suit, rank, and color
type Card struct {
	suit  Suit
	rank  Rank
	color Color
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.rank, strings.ToTitle(c.suit))
}

package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

// Color of a card aliased to string
type Color = string

// Suit of a card aliased to string
type Suit = string

// Suits of cards
var (
	Spades   Suit = "Spades"
	Clubs    Suit = "Clubs"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
)

// Colors of cards
const (
	Black Color = "black"
	Red   Color = "red"
)

var suitToColor = map[Suit]Color{
	Spades:   Black,
	Clubs:    Black,
	Hearts:   Red,
	Diamonds: Red,
}

var suitToValue = map[Suit]int{
	Spades:   1,
	Diamonds: 2,
	Clubs:    3,
	Hearts:   4,
}

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
	Joker Rank = 0
)

var rankToHumanString = map[Rank]string{
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
	Joker: "Joker",
}

func (r Rank) String() string {
	return rankToHumanString[r]
}

// Card represents a normal playing card with a suit, rank, and color
type Card struct {
	Suit  Suit
	Rank  Rank
	Color Color
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

// New creates a new deck
func New(opts ...func([]Card) []Card) []Card {
	deck := createDeck()
	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck
}

func createDeck() []Card {
	deckLength := 52
	deck := make([]Card, 0, deckLength)
	suits := []Suit{Spades, Diamonds, Clubs, Hearts}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	for _, s := range suits {
		for _, r := range ranks {
			deck = append(deck, Card{
				Suit:  s,
				Rank:  r,
				Color: suitToColor[s],
			})
		}
	}
	return deck
}

// Sort a deck of cards with the provided comparison function
func Sort(compare func(deck []Card) func(i, j int) bool) func(deck []Card) []Card {
	return func(deck []Card) []Card {
		sort.Slice(deck, compare(deck))
		return deck
	}
}

// Filter a deck of cards with the provided filter function
// filter should return true to filter an element out
func Filter(filter func(c Card) bool) func(deck []Card) []Card {
	return func(deck []Card) []Card {
		filteredCards := make([]Card, 0, len(deck))
		for _, c := range deck {
			if !filter(c) {
				filteredCards = append(filteredCards, c)
			}
		}
		return filteredCards
	}
}

// Shuffle a deck of cards in random order
func Shuffle(deck []Card) []Card {
	shuffledCards := make([]Card, len(deck))

	positions := rand.Perm(len(deck))

	for i, j := range positions {
		shuffledCards[i] = (deck)[j]
	}

	return shuffledCards
}

// NewOrder provides a less comparison to put a deck in an
// ordering as if the deck was new
func NewOrder(deck []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(deck[i]) < absRank(deck[j])
	}
}

func absRank(c Card) int {
	return int(c.Rank) * int(suitToValue[c.Suit])
}

// AddJokers adds n jokers to the end of the deck
func AddJokers(n uint) func(deck []Card) []Card {
	return func(deck []Card) []Card {
		for i := uint(0); i < n; i++ {
			deck = append(deck, Card{Rank: Joker, Color: "", Suit: ""})
		}
		return deck
	}
}

// MultipleDecks adds n decks together to form a single deck
func MultipleDecks(n uint) func(deck []Card) []Card {
	return func(deck []Card) []Card {
		for i := uint(0); i < n-1; i++ {
			deck = append(deck, createDeck()...)
		}
		return deck
	}
}

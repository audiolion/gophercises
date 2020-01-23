package deck

import (
	"reflect"
	"testing"
)

func TestCard(t *testing.T) {
	card := Card{
		suit:  Spades,
		color: Black,
		rank:  Ace,
	}

	got := card.String()
	want := "Ace of Spades"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNew(t *testing.T) {
	t.Run("make a plain new deck", func(t *testing.T) {
		deck := New()

		got := len(deck)
		want := 52

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

		deckSuitRangesTest := []struct {
			desc  string
			start int
			end   int
			want  Suit
		}{
			{desc: "First 13 should be Spades", start: 0, end: 12, want: Spades},
			{desc: "Second 13 should be Diamonds", start: 13, end: 25, want: Diamonds},
			{desc: "Third 13 should be Clubs", start: 26, end: 38, want: Clubs},
			{desc: "Fourth 13 should be Hearts", start: 39, end: 52, want: Hearts},
		}

		for _, tt := range deckSuitRangesTest {
			for i := tt.start; i < tt.end; i++ {
				t.Run(tt.desc, func(t *testing.T) {
					got := deck[i].suit
					if got != tt.want {
						t.Errorf("%#v got %q want %q", deck[i], got, tt.want)
					}
				})
			}
		}
	})

	t.Run("make a new deck that is shuffled", func(t *testing.T) {
		deck := New(Shuffle)

		if deck[0].suit == Spades && deck[0].rank == Ace {
			t.Errorf("expected new deck to be shuffled")
		}
	})

	t.Run("make a new deck with jokers", func(t *testing.T) {
		deck := New(AddJokers(3))

		for _, c := range deck[len(deck)-3:] {
			if c.rank != Joker {
				t.Errorf("got %q want %q", c.rank, Joker)
			}
		}
	})

	t.Run("custom filter of all twos and threes from a deck", func(t *testing.T) {
		filterTwosThrees := func(c Card) bool {
			return c.rank == Two || c.rank == Three
		}

		deck := New(Filter(filterTwosThrees))

		for _, c := range deck {
			if c.rank == Two || c.rank == Three {
				t.Errorf("expected no twos or threes in deck")
			}
		}
	})

	t.Run("make a new deck of multiple decks", func(t *testing.T) {
		deck := New(MultipleDecks(3))

		got := len(deck)
		want := 52 * 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSort(t *testing.T) {
	t.Run("custom compare", func(t *testing.T) {
		deck := New()
		deck[0], deck[1] = deck[1], deck[0]

		compare := func(cards []Card) func(i, j int) bool {
			return func(i, j int) bool {
				return cards[i].rank > cards[j].rank
			}
		}

		deck = Sort(compare)(deck)

		got := deck[0].rank
		want := King
		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	})

	t.Run("default sort", func(t *testing.T) {
		deck := New()

		deck = Shuffle(deck)

		deck = Sort(NewOrder)(deck)

		got := deck[0].rank
		want := Ace
		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
		got2 := deck[0].suit
		want2 := Spades
		if got2 != want2 {
			t.Errorf("got %q expected %q", got2, want2)
		}
	})
}

func TestShuffle(t *testing.T) {
	deck := New()

	shuffledDeck := Shuffle(deck)

	if reflect.DeepEqual(shuffledDeck, deck) {
		t.Errorf("Cards were not shuffled")
	}
}

package deck

import "testing"

func TestCard(t *testing.T) {
	card := Card{
		suit:  Spades,
		color: Black,
		rank:  1,
	}

	got := card.String()
	want := "Ace of Spades"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

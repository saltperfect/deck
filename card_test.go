package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Heart})
	fmt.Println(Card{Rank: Seven, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Hearts
	// Seven of Hearts
	// Jack of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Not corect number of cards")
	}
}

func TestDefautSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Suit: Spade, Rank: Ace}
	if cards[0] != exp {
		t.Error("first card was not ace of spades")
	}
}
func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Suit: Spade, Rank: Ace}
	if cards[0] != exp {
		t.Error("first card was not ace of spades")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker{
			count++
		}
	}
	if count != 3 {
		t.Error("Correct number of jokers not added")
	}
}

func TestFilter(t *testing.T) {
	filter := func( card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Filter failed")
		}
	}
}
func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("want %d, got %d ", 13*4*3, len(cards))
	}
}

func TestSuffle(t *testing.T) {
	// overriding suffle rank value to make things deterministic
	// [40 35 50 0 44 7 1 16 13 4 21 12 23 34 19 11 42 20 17 48 27 9 43 46 47 45 5 49 51 30 41 26 25 32 39 28 37 31 33 10 22 8 6 29 36 18 14 2 15 3 38 24]
	suffleRand = rand.New(rand.NewSource(0))

	orig := New()
	suffled := New(Suffle)
	first := orig[40]
	second := orig[35]

	if suffled[0] != first {
		t.Errorf( " Expect first to be %s, got %s", first, suffled[0])
	}
	if suffled[1] != second {
		t.Errorf( " Expect second to be %s, got %s", second, suffled[1])
	}

}
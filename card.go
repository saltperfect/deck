//go:generate stringer -type=Suit,Rank
// Package deck provides you a way to generate deck of cards by the help of functional options
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit enum represents cards suit
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank enum represents cards rank
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card type represents card type
type Card struct {
	Suit
	Rank
}

// String funtion for stringer support
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New takes in funtional options as variadic parameters and return a slice of cards
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// DefaultSort is a functional option that sorts deck via suit and then rank
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort function is a custom sort which takes in a function which takes slice of card and returns
// function(i, j int) bool determining the logic of sorting
// Take a look at Less function below
func Sort(less func([]Card) func(i, j int) bool) func([]Card) []Card {
	return func(c []Card) []Card {
		sort.Slice(c, less(c))
		return c
	}
}

// Less function that is getting used in default sort. take this function as a example for the type function need to be passed in Sort
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var suffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Suffle funtional optional provides a suffled card
func Suffle(cards []Card) []Card {
	ret := make([]Card, len(cards))

	for i, j := range suffleRand.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}

// Use to add n number of Jokers
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		jslice := make([]Card, n)
		for i := range jslice {
			jslice[i] = Card{Suit: Joker, Rank: Rank(i)}
		}
		return append(cards, jslice...)
	}
}

// Filter out cards before generating the new deck
// takes in a function(Card) bool and return a funtional option
func Filter(f func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// Deck takes in number of decks that needs to be generated
func Deck(n int) func([]Card) []Card {
	return func(c []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, c...)
		}
		return ret
	}
}
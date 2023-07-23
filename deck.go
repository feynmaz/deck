package deck

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
)

// Deck represents a deck of cards, which is a slice of type Card
type Deck []Card

// New creates new Deck
func New(count int, filter ...Rank) Deck {
	deck := make([]Card, 0, 52)

	for _, rank := range Ranks {
		if slices.Contains(filter, rank) {
			continue
		}
		for _, suit := range Suits {
			card := Card{
				Suit: suit,
				Rank: rank,
			}
			deck = append(deck, card)
		}
	}

	if count > 1 {
		result := make(Deck, 0, len(deck)*count)
		for i := 0; i < count; i++ {
			result = append(result, deck...)
		}
		return result
	}
	return deck
}

// Shuffle shuffles deck of cards in random order
func (deck Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	randGenerator.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

// Len implements Len() of sort.Interface
func (deck Deck) Len() int {
	return len(deck)
}

// Len implements Less() of sort.Interface
func (deck Deck) Less(i, j int) bool {
	return deck[i].GetAbsRank() < deck[j].GetAbsRank()
}

// Len implements Swap() of sort.Interface
func (deck Deck) Swap(i, j int) {
	deck[i], deck[j] = deck[j], deck[i]
}

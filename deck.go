package deck

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
)

type Deck []Card

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

func (deck Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	randGenerator.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

// Implementing sort.Interface https://pkg.go.dev/sort#Interface
func (deck Deck) Len() int {
	return len(deck)
}

func (deck Deck) Less(i, j int) bool {
	return deck[i].GetAbsRank() < deck[j].GetAbsRank()
}

func (deck Deck) Swap(i, j int) {
	deck[i], deck[j] = deck[j], deck[i]
}

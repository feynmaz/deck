package deck

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	deck := New(1)
	deckFiltered := New(1, Two, Three, Four, Five)

	require.True(t, len(deck) > len(deckFiltered))

	require.Equal(t, Two, deck[4].Rank)
	require.NotEqual(t, Two, deckFiltered[4].Rank)

	deck = New(4)
	require.Equal(t, len(deck), 52*4)
}

func Test_DefaultSort(t *testing.T) {
	deck := New(1)

	deckCopy := make(Deck, 0, len(deck))
	deckCopy = append(deckCopy, deck...)

	deckCopy[51], deckCopy[4] = deckCopy[4], deckCopy[51]
	deckCopy[49], deckCopy[8] = deckCopy[8], deckCopy[49]
	require.NotEqual(t, deck, deckCopy)
	require.NotEqual(t, deckCopy[51].Rank, King)

	sort.Sort(deckCopy)
	require.Equal(t, deckCopy[48].Rank, King)
	require.Equal(t, deckCopy[49].Rank, King)
	require.Equal(t, deckCopy[50].Rank, King)
	require.Equal(t, deckCopy[51].Rank, King)
}

func Test_Shuffle(t *testing.T) {
	deck := New(1)
	deckCopy := New(1)

	deck.Shuffle()
	require.NotEqual(t, deck, deckCopy)
}

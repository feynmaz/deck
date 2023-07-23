package deck

// go:generate stringer -type=Suit,Rank

// Suit represents card suit
type Suit uint8

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
)

var Suits = []Suit{Spades, Diamonds, Clubs, Hearts}

// Rank represents card rank
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

var Ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Card represents card
type Card struct {
	Suit
	Rank
}

// String implements Stringer interface
func (c Card) String() string {
	return c.Rank.String() + " of " + c.Suit.String()
}

// GetAbsRank returns card rank to be used in compare functions
func (c Card) GetAbsRank() int {
	return int(c.Rank)
}

package deck

// go:generate stringer -type=Suit,Rank

type Suit uint8

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
)

var Suits = []Suit{Spades, Diamonds, Clubs, Hearts}

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

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return c.Rank.String() + " of " + c.Suit.String()
}

func (c Card) GetAbsRank() int {
	return int(c.Rank)
}

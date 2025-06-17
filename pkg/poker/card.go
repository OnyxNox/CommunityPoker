package poker

import "math/rand"

// Represents a playing card suit.
type Suit uint8

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

// Converts the enumerator's numeric value to its corresponding string.
func (suit Suit) String() string {
	return [...]string{"♣", "♦", "♥", "♠"}[suit]
}

// Converts the enumerator's numeric value to its corresponding string during JSON encoding.
func (suit Suit) MarshalJSON() ([]byte, error) {
	return []byte(`"` + suit.String() + `"`), nil
}

// Represent a playing card.
type Card struct {
	// Playing card rank. A = 1, J = 11, Q = 12, K = 13
	Rank uint8

	// Playing card suit.
	Suit Suit
}

// Initializes a new deck of playing cards.
func NewDeck() []Card {
	var deck []Card
	for suit := Clubs; suit <= Spades; suit++ {
		for rank := uint8(1); rank <= 13; rank++ {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}

	return deck
}

// Shuffles a deck of playing cards randomly.
func Shuffle(deck []Card) {
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)

		deck[i], deck[j] = deck[j], deck[i]
	}
}

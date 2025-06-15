package poker

// Represents a playing card suit.
type Suit uint8

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (suit Suit) ToString(toSymbol bool) string {
	var outputStrings [4]string
	if toSymbol {
		outputStrings = [...]string{"♣", "♦", "♥", "♠"}
	} else {
		outputStrings = [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	}

	return outputStrings[suit]
}

// Represent a playing card.
type Card struct {
	// Playing card rank. A = 1, J = 11, Q = 12, K = 13
	Rank uint8

	// Playing card suit.
	Suit Suit
}

func NewDeck() []Card {
	var deck []Card
	for suit := Clubs; suit <= Spades; suit++ {
		for rank := uint8(1); rank <= 13; rank++ {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}

	return deck
}

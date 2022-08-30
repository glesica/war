package sim

import (
	"math/rand"
	"time"
)

type Card uint8

type Deck []Card

func ShuffledDeck() Deck {
	rand.Seed(time.Now().UnixNano())

	deck := OrderedDeck()
	Shuffle(deck)

	return deck
}

func OrderedDeck() Deck {
	var deck Deck
	for suit := 0; suit < 4; suit++ {
		for value := 2; value <= 14; value++ {
			deck = append(deck, Card(value))
		}
	}

	return deck
}

func Deal(deck Deck) (deck1 Deck, deck2 Deck) {
	// Weird artifact of shuffling, maybe, the bottom of the deck
	// seems to have more high cards. So instead of just cutting it,
	// we deal the way humans would.
	deal1 := true
	for _, card := range deck {
		if deal1 {
			deck1 = append(deck1, card)
		} else {
			deck2 = append(deck2, card)
		}
		deal1 = !deal1
	}
	return
}

func Draw(deck Deck, count int) (drawn Deck, updated Deck) {
	if len(deck) < count {
		return deck[:], nil
	}

	return deck[:count], deck[count:]
}

func Empty(deck Deck) bool {
	return len(deck) == 0
}

func Shuffle(deck Deck) {
	rand.Seed(time.Now().UnixNano())

	count := len(deck)
	for first := 0; first < count; first++ {
		next := rand.Intn(count-first) + first
		deck[first], deck[next] = deck[next], deck[first]
	}
}

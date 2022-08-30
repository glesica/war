package sim

import "math"

func count(deck Deck) int {
	return len(deck)
}

func minCount(deck1, deck2 Deck) int {
	return int(math.Min(float64(count(deck1)), float64(count(deck2))))
}

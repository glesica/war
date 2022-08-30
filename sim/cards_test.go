package sim

import "testing"

func TestDeal(t *testing.T) {
	h1, h2 := Deal([]Card{2, 3, 4, 5})
	assertDecksMatch(t, h1, []Card{2, 4})
	assertDecksMatch(t, h2, []Card{3, 5})
}

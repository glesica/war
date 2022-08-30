package sim

import (
	"strconv"
	"testing"
)

func TestPlay(t *testing.T) {
	for _, c := range []struct {
		start1 Deck
		start2 Deck
		end1   Deck
		end2   Deck
	}{
		{[]Card{5}, []Card{4}, []Card{5, 4}, []Card{}},
		{[]Card{4}, []Card{5}, []Card{}, []Card{4, 5}},
		{[]Card{5, 6}, []Card{5, 4}, []Card{5, 5, 6, 4}, []Card{}},
		{[]Card{5, 2, 6}, []Card{5, 3, 4}, []Card{5, 5, 2, 3, 6, 4}, []Card{}},
	} {
		h1, h2 := Play(c.start1, c.start2)
		assertCardsMatch(t, h1, c.end1)
		assertCardsMatch(t, h2, c.end2)
	}
}

func assertCardsMatch(t *testing.T, actual, expected Deck) {
	format := "%s:\nactual: %v\nexpected: missing"

	for _, actualCard := range actual {
		match := false
		for _, expectedCard := range expected {
			match = actualCard == expectedCard
			if match {
				break
			}
		}
		if !match {
			t.Errorf(format, "unexpected card", actualCard)
		}
	}

	format = "%s:\nactual: missing\nexpected: %v"

	for _, expectedCard := range expected {
		match := false
		for _, actualCard := range actual {
			match = expectedCard == actualCard
			if match {
				break
			}
		}
		if !match {
			t.Errorf(format, "missing card", expectedCard)
		}
	}
}

func assertDecksMatch(t *testing.T, actual, expected Deck) {
	format := "%s:\nactual: %v\nexpected: %v"

	if len(actual) != len(expected) {
		t.Errorf(format, "decks are different length", actual, expected)
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf(format, "mismatch at position "+strconv.Itoa(i), actual, expected)
		}
	}
}

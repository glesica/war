package sim

func Full(in1, in2 Deck, maxRounds int) (Deck, Deck, int, bool) {
	rounds := 0
	for !Over(in1, in2) {
		rounds++
		in1, in2 = Play(in1, in2)

		if rounds >= maxRounds {
			return in1, in2, rounds, false
		}
	}

	return in1, in2, rounds, true
}

func Play(in1, in2 Deck) (out1, out2 Deck) {
	if Empty(in1) || Empty(in2) {
		panic("invariant: one deck is already empty")
	}

	cards1, in1 := Draw(in1, 1)
	cards2, in2 := Draw(in2, 1)

	var pot Deck

	for cards1[0] == cards2[0] {
		pot = append(pot, cards1...)
		pot = append(pot, cards2...)

		// This should almost never happen, but if both players
		// have identical hands we can end up with no winner.
		if count(in1) == 0 && count(in2) == 0 {
			return Deal(ShuffledDeck())
		}

		// Player 1 is out of cards and can't break the tie,
		// so player 2 wins.
		if count(in1) == 0 {
			return nil, append(in2, pot...)
		}

		// Player 2 is out of cards and can't break the tie,
		// so player 1 wins.
		if count(in2) == 0 {
			return append(in1, pot...), nil
		}

		// Figure out how many cards, up to three, we can
		// burn such that both players will still have at
		// least one card to play.
		burnCount := minCount(in1, in2) - 1
		if burnCount > 3 {
			burnCount = 3
		}

		// Burn cards by adding them to the pot.
		cards1, in1 = Draw(in1, burnCount)
		cards2, in2 = Draw(in2, burnCount)
		pot = append(pot, cards1...)
		pot = append(pot, cards2...)

		// Draw another card to determine who gets the pot.
		cards1, in1 = Draw(in1, 1)
		cards2, in2 = Draw(in2, 1)
	}

	pot = append(pot, cards1...)
	pot = append(pot, cards2...)

	// Shuffle the pot to avoid pathological periodic behavior.
	// This causes more (all?) games to eventually end and more
	// realistically simulates how humans would play the game.
	Shuffle(pot)

	if cards1[0] > cards2[0] {
		return append(in1, pot...), in2
	}

	if cards2[0] > cards1[0] {
		return in1, append(in2, pot...)
	}

	panic("invariant: hand had no winner")
}

func Over(in1, in2 Deck) bool {
	win1 := Empty(in2)
	win2 := Empty(in1)
	return win1 || win2
}

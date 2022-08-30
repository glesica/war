package main

import (
	"flag"
	"fmt"
	"github.com/glesica/war/analyze"
	"github.com/glesica/war/sim"
	"os"
)

func main() {
	hist := analyze.NewHist(30)

	roundsCap := flag.Int("rounds-cap", 1000, "max number of rounds per game")
	totalGames := flag.Int("games", 100, "number of games to play")
	flag.Parse()

	for i := 0; i < *totalGames; i++ {
		_, _ = fmt.Fprintf(os.Stderr, "playing game %d of %d\n", i+1, *totalGames)

		hand1, hand2 := sim.Deal(sim.ShuffledDeck())
		hand1, hand2, rounds, finished := sim.Full(hand1, hand2, *roundsCap)

		if finished {
			hist.Add(rounds)
		}
	}

	hist.Print(70)
}

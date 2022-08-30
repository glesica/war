# War

I wanted to find out how long the card game War should be expected to take with
two players.

The histogram below is the result of 50,000 simulated games. The vertical axis
is the number of turns the games took, the horizontal axis is the number of
games that took that many turns (bucketed).

Interestingly, only about 35% of games take fewer than 150 turns, which is about
as long as most people probably allow the game to go on before quitting (at
least in my mind). That means that 65% of games are likely to be abandoned!

```
0000 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 11.70%
0074 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 24.39%
0148 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 18.78%
0222 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 13.26%
0296 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 9.34%
0370 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 6.70%
0444 | ◼◼◼◼◼◼◼◼◼◼◼◼◼ 4.73%
0518 | ◼◼◼◼◼◼◼◼◼ 3.19%
0592 | ◼◼◼◼◼◼ 2.34%
0666 | ◼◼◼◼ 1.66%
0740 | ◼◼◼ 1.14%
0814 | ◼◼ 0.81%
0888 | ◼ 0.57%
0962 | ◼ 0.43%
1036 | ◼ 0.30%
1110 | ◼ 0.18%
1184 | ◼ 0.13%
1258 | ◼ 0.12%
1332 | ◼ 0.07%
1406 | ◼ 0.04%
1480 | ◼ 0.04%
1554 | ◼ 0.03%
1628 | ◼ 0.02%
1702 | ◼ 0.01%
1776 | ◼ 0.00%
1850 | ◼ 0.01%
1924 | ◼ 0.00%
1998 | ◼ 0.01%
2072 |
2146 | ◼ 0.00%
```

Below is another histogram from a simulation of 100,000 games. Note that
the general shape is the same.

```
0000 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 20.47%
0100 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 29.51%
0200 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 18.83%
0300 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 11.75%
0400 | ◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼◼ 7.35%
0500 | ◼◼◼◼◼◼◼◼◼◼ 4.50%
0600 | ◼◼◼◼◼◼ 2.89%
0700 | ◼◼◼◼ 1.74%
0800 | ◼◼ 1.11%
0900 | ◼ 0.68%
1000 | ◼ 0.45%
1100 | ◼ 0.27%
1200 | ◼ 0.19%
1300 | ◼ 0.10%
1400 | ◼ 0.06%
1500 | ◼ 0.05%
1600 | ◼ 0.02%
1700 | ◼ 0.01%
1800 | ◼ 0.01%
1900 | ◼ 0.00%
2000 | ◼ 0.00%
2100 | ◼ 0.00%
2200 | ◼ 0.00%
2300 |
2400 |
2500 |
2600 |
2700 |
2800 |
2900 | ◼ 0.00%
```

## Run

To run, build the binary with `go build -o war .`. Then, run `./war` with
default options for a small simulation. See `-help` for available parameters.

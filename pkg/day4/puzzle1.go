package day4

import (
	"math"
	"slices"
)

func Puzzle1(input string) int {
	points := 0
	games := New(input)

	for _, game := range games {
		match := 0.0
		for _, available := range game.Available {
			if slices.Contains(game.Winning, available) {
				match++
			}
		}

		points += int(math.Pow(2, match-1))
	}

	return points
}

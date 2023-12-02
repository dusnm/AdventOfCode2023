package day2

import (
	"strings"
)

func Puzzle1(input string) int {
	sum := 0
	totalRed := 12
	totalGreen := 13
	totalBlue := 14

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		id, gamePart := tokenizeLine(line)

		gameIsValid := true
		for _, game := range gamePart {
			gameListing := strings.Split(game, ", ")
			currentRed, currentGreen, currentBlue := countCubes(gameListing)

			if currentRed > totalRed || currentGreen > totalGreen || currentBlue > totalBlue {
				gameIsValid = false
			}
		}

		if gameIsValid {
			sum += id
		}
	}

	return sum
}

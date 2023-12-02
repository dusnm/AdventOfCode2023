package day2

import (
	"strings"
)

func Puzzle2(input string) int {
	sumOfPowers := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		_, gamePart := tokenizeLine(line)
		for _, game := range gamePart {
			gameListing := strings.Split(game, ", ")
			currentRed, currentGreen, currentBlue := countCubes(gameListing)

			if currentRed > maxRed {
				maxRed = currentRed
			}

			if currentGreen > maxGreen {
				maxGreen = currentGreen
			}

			if currentBlue > maxBlue {
				maxBlue = currentBlue
			}
		}

		sumOfPowers += maxRed * maxGreen * maxBlue
	}

	return sumOfPowers
}

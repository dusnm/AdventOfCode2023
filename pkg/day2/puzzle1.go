package day2

import (
	"strconv"
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

		tokens := strings.FieldsFunc(line, func(r rune) bool {
			return r == ':' || r == ';'
		})

		idPart := strings.Replace(tokens[0], "Game ", "", 1)
		gamePart := tokens[1:]

		// This is fine 🔥🐶🔥
		id, _ := strconv.Atoi(idPart)

		gameIsValid := true
		for _, game := range gamePart {
			parts := strings.Split(game, ", ")
			for _, part := range parts {
				currentRed := 0
				currentGreen := 0
				currentBlue := 0

				part = strings.TrimSpace(part)

				// Exploiting the fact that the number of cubes
				// is at most two digits long
				value := strings.TrimSpace(part[0:2])

				// Moving the index of the first letter of the color accordingly
				var color rune
				if len(value) > 1 {
					color = rune(part[3])
				} else {
					color = rune(part[2])
				}

				// This is fine 🔥🐶🔥
				v, _ := strconv.Atoi(value)
				switch color {
				case 'r':
					currentRed += v
				case 'g':
					currentGreen += v
				case 'b':
					currentBlue += v
				}

				if currentRed > totalRed || currentGreen > totalGreen || currentBlue > totalBlue {
					gameIsValid = false
				}
			}
		}

		if gameIsValid {
			sum += id
		}
	}

	return sum
}

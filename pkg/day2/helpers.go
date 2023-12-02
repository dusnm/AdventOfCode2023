package day2

import (
	"strconv"
	"strings"
)

func tokenizeLine(line string) (int, []string) {
	tokens := strings.FieldsFunc(line, func(r rune) bool {
		return r == ':' || r == ';'
	})

	idstr := strings.Replace(tokens[0], "Game ", "", 1)

	// This is fine 🔥🐶🔥
	id, _ := strconv.Atoi(idstr)

	return id, tokens[1:]
}

func countCubes(gameListing []string) (int, int, int) {
	cubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, cube := range gameListing {
		cube = strings.TrimSpace(cube)
		values := strings.Split(cube, " ")

		// This is fine 🔥🐶🔥
		value, _ := strconv.Atoi(values[0])
		color := values[1]

		cubes[color] = value
	}

	return cubes["red"], cubes["green"], cubes["blue"]
}

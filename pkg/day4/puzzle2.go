package day4

import (
	"slices"
)

func Puzzle2(input string) int {
	games := New(input)
	gameMap := make(map[int]Game, len(games))
	for _, game := range games {
		gameMap[game.ID] = game
	}

	// Range based for loop is impossible
	// since the slice is modified in-place
	for i := 0; i < len(games); i++ {
		id := games[i].ID
		for _, available := range games[i].Available {
			if slices.Contains(games[i].Winning, available) {
				id++
				games = append(games, gameMap[id])
			}
		}
	}

	return len(games)
}

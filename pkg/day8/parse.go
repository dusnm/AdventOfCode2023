package day8

import (
	"strings"
)

func New(input string) Game {
	game := Game{}
	game.Nodes = make(map[string]Node)

	lines := strings.Split(input, "\n")
	instructions, maps := lines[0], lines[1:]

	for _, c := range instructions {
		game.Instructions = append(game.Instructions, c)
	}

	for _, line := range maps {
		if len(line) == 0 {
			continue
		}

		replacer := strings.NewReplacer(
			" = (", ",",
			", ", ",",
			")", "",
		)

		line = replacer.Replace(line)
		data := strings.Split(line, ",")

		node := Node{}
		node.Value = data[0]
		node.Left = data[1]
		node.Right = data[2]

		game.Nodes[node.Value] = node
	}

	return game
}

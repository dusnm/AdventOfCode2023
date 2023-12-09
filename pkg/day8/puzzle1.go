package day8

const (
	GameStart = "AAA"
	GameEnd   = "ZZZ"
)

func Puzzle1(input string) int {
	game := New(input)

	depth, next := 0, 0
	lookup := GameStart

	for {
		var instruction rune
		if len(game.Instructions) > next {
			instruction = game.Instructions[next]
		} else {
			instruction = game.Instructions[0]
			next = 0
		}

		value := game.Nodes[lookup]

		if value.Value == GameEnd {
			break
		}

		if instruction == 'L' {
			lookup = value.Left
		} else {
			lookup = value.Right
		}

		next++
		depth++
	}

	return depth
}

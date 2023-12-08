package day8

const (
	GameStart            = "AAA"
	GameEnd              = "ZZZ"
	InstructionLeft rune = 'L'
)

func Puzzle1(input string) int {
	game := New(input)

	return Walk(&game, GameStart, 0, false, 0)
}

func Walk(
	game *Game,
	lookup string,
	nextInstruction int,
	reachedEnd bool,
	depth int,
) int {
	if reachedEnd {
		return depth
	}

	var instruction rune

	if len(game.Instructions) > nextInstruction {
		instruction = game.Instructions[nextInstruction]
	} else {
		instruction = game.Instructions[0]
		nextInstruction = 0
	}

	value := game.Nodes[lookup]

	if value.Value == GameEnd {
		return Walk(game, value.Value, nextInstruction, true, depth)
	}

	nextInstruction++
	depth++

	if instruction == InstructionLeft {
		return Walk(game, value.Left, nextInstruction, false, depth)
	}

	return Walk(game, value.Right, nextInstruction, false, depth)
}

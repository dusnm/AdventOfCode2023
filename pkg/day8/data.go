package day8

type (
	Node struct {
		Left  string
		Right string
		Value string
	}

	Game struct {
		Instructions []rune
		Nodes        map[string]Node
	}
)

func (g *Game) Walk(
	lookup string,
	nextInstruction int,
	reachedEnd bool,
	depth int,
) int {
	if reachedEnd {
		return depth
	}

	var instruction rune

	if len(g.Instructions) > nextInstruction {
		instruction = g.Instructions[nextInstruction]
	} else {
		instruction = g.Instructions[0]
		nextInstruction = 0
	}

	value := g.Nodes[lookup]

	if value.Value == "ZZZ" {
		return g.Walk(value.Value, nextInstruction, true, depth)
	}

	nextInstruction++
	depth++

	if instruction == 'L' {
		return g.Walk(value.Left, nextInstruction, false, depth)
	}

	return g.Walk(value.Right, nextInstruction, false, depth)
}

package day8

import (
	"sync"
)

func Puzzle2(input string) int {
	game := New(input)

	startingNodes := make([]string, 0, len(game.Nodes))
	for k := range game.Nodes {
		if k[len(k)-1] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}

	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	depths := make([]int, 0, len(startingNodes))

	wg.Add(len(startingNodes))
	for _, node := range startingNodes {
		go func(
			wg *sync.WaitGroup,
			mu *sync.RWMutex,
			node string,
			instructions []rune,
			nodes map[string]Node,
		) {
			defer wg.Done()

			depth, next := 0, 0
			lookup := node

			for {
				var instruction rune
				if len(game.Instructions) > next {
					instruction = game.Instructions[next]
				} else {
					instruction = game.Instructions[0]
					next = 0
				}

				value := game.Nodes[lookup]

				if value.Value[len(value.Value)-1] == 'Z' {
					mu.Lock()
					depths = append(depths, depth)
					mu.Unlock()

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
		}(
			wg,
			mu,
			node,
			game.Instructions,
			game.Nodes,
		)
	}

	wg.Wait()

	answer, _ := lcm(depths...)

	return answer
}

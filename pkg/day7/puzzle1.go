package day7

import "fmt"

func Puzzle1(input string) int {
	game := New(input)
	game.Sort()

	total, rank := 0, 1
	for _, hand := range game.Hands {
		fmt.Println(hand.Value, rank)
		total += hand.Bid * rank

		rank++
	}

	return total
}

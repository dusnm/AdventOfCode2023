package day7

func Puzzle2(input string) int {
	withJokers := true
	game := New(input)

	game.Sort(withJokers)

	total, rank := 0, 1
	for _, hand := range game.Hands {
		total += hand.Bid * rank

		rank++
	}

	return total
}

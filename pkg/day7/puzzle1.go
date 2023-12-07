package day7

func Puzzle1(input string) int {
	withJokers := false
	game := New(input)

	game.Sort(withJokers)

	total, rank := 0, 1
	for _, hand := range game.Hands {
		total += hand.Bid * rank

		rank++
	}

	return total
}

package day9

func Puzzle2(input string) int {
	numbers := New(input)

	total := 0

	for _, report := range numbers {
		total += extrapolateBackwards(report)
	}

	return total
}

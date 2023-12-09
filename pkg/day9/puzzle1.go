package day9

func Puzzle1(input string) int {
	numbers := New(input)

	total := 0
	for _, report := range numbers {
		total += extrapolateForwards(report)
	}

	return total
}

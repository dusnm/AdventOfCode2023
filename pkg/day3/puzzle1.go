package day3

type (
	PartNumber struct {
		Value     int
		Length    int
		Adjacency []Adjacent
	}
)

func Puzzle1(input string) int {
	sumOfPartNumbers := 0
	partNumbers := findPartNumbers(load(input))

	for _, part := range partNumbers {
		sumOfPartNumbers += part.Value
	}

	return sumOfPartNumbers
}

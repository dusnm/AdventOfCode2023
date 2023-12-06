package day6

func Puzzle2(input string) int {
	r := NewSingleRace(input)

	winCounter := 0
	for holdLength := 0; holdLength < r.Time; holdLength++ {
		if r.IsEnoughToBeatRecord(holdLength) {
			winCounter++
		}
	}

	return winCounter
}

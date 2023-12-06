package day6

func Puzzle1(input string) int {
	races := NewMultipleRaces(input)
	totalWins := 1

	for _, r := range races {
		winCounter := 0
		for holdLength := 0; holdLength < r.Time; holdLength++ {
			if r.IsEnoughToBeatRecord(holdLength) {
				winCounter++
			}
		}

		totalWins *= winCounter
	}

	return totalWins
}

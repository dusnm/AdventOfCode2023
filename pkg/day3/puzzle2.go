package day3

func Puzzle2(input string) int {
	sumOfGearRatios := 0
	partNumbers := findPartNumbers(load(input))

	gears := map[string][]int{}
	for _, partNumber := range partNumbers {
		for _, adjacency := range partNumber.Adjacency {
			if adjacency.Value == '*' {
				gearPartNumbers := gears[adjacency.String()]
				gearPartNumbers = append(gearPartNumbers, partNumber.Value)
				gears[adjacency.String()] = gearPartNumbers
				break
			}
		}
	}

	for _, v := range gears {
		if len(v) == 2 {
			sumOfGearRatios += v[0] * v[1]
		}
	}

	return sumOfGearRatios
}

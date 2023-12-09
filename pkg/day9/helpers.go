package day9

func generateSequence(report []int) []int {
	sequence := make([]int, 0, len(report)-1)

	for i := 1; i < len(report); i++ {
		current := report[i]
		previous := report[i-1]

		sequence = append(sequence, current-previous)
	}

	return sequence
}

func extrapolateForwards(report []int) int {
	values := make([]int, 0, len(report))

	for {
		allZero := true
		values = append(values, report[len(report)-1])
		sequence := generateSequence(report)

		for _, v := range sequence {
			allZero = allZero && v == 0
		}

		if !allZero {
			report = sequence
		} else {
			break
		}
	}

	extrapolated := 0
	for _, v := range values {
		extrapolated += v
	}

	return extrapolated
}

func extrapolateBackwards(report []int) int {
	values := make([]int, 0, len(report))

	for {
		allZero := true
		values = append(values, report[0])
		sequence := generateSequence(report)

		for _, v := range sequence {
			allZero = allZero && v == 0
		}

		if !allZero {
			report = sequence
		} else {
			break
		}
	}

	extrapolated := 0
	for i := len(values) - 1; i >= 0; i-- {
		extrapolated = values[i] - extrapolated
	}

	return extrapolated
}

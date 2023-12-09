package day8

import "errors"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(numbers ...int) (int, error) {
	if len(numbers) < 2 {
		return 0, errors.New("at least two numbers requiered for calculation")
	}

	a, b := numbers[0], numbers[1]
	result := a * b / gcd(a, b)

	for i := 2; i < len(numbers); i++ {
		result = result * numbers[i] / gcd(result, numbers[i])
	}

	return result, nil
}

package day7

import (
	"strconv"
	"strings"
)

func New(input string) Game {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		data := strings.Split(line, " ")

		bid, _ := strconv.Atoi(data[1])

		hand := Hand{
			Value: data[0],
			Bid:   bid,
		}

		hands = append(hands, hand)
	}

	return Game{Hands: hands}
}

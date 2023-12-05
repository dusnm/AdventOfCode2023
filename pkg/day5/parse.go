package day5

import (
	"bufio"
	"strings"
	"unicode"
)

const (
	StateSeeds int = iota
	StateSeedToSoil
	StateSoilToFertilizer
	StateFertilizerToWater
	StateWaterToLight
	StateLightToTemperature
	StateTemperatureToHumidity
	StateHumidityToLocation
)

type (
	transition struct {
		state []rune
	}
)

func (t *transition) Append(c rune) {
	t.state = append(t.state, c)
}

func (t *transition) Next(c rune) bool {
	if len(t.state) == 0 {
		return false
	}

	return t.state[len(t.state)-1] == c
}

func (t *transition) Reset() {
	t.state = make([]rune, 0)
}

func addToRange(current *[]Range, c rune, index int, key int) {
	if len(*current) <= index {
		r := Range{}

		switch key {
		case RangeDestination:
			r.Destination = int(c - '0')
		case RangeSource:
			r.Source = int(c - '0')
		case RangeLength:
			r.Length = int(c - '0')
		}

		*current = append(*current, r)
	} else {
		switch key {
		case RangeDestination:
			data := *current
			value := data[index].Destination*10 + int(c-'0')
			data[index].Destination = value
			*current = data
		case RangeSource:
			data := *current
			value := data[index].Source*10 + int(c-'0')
			data[index].Source = value
			*current = data
		case RangeLength:
			data := *current
			value := data[index].Length*10 + int(c-'0')
			data[index].Length = value
			*current = data
		}
	}
}

func parseRange(
	c rune,
	_range *[]Range,
	index *int,
	key *int,
	t *transition,
) bool {
	if unicode.IsDigit(c) {
		addToRange(_range, c, *index, *key)
	}

	if c == ' ' {
		*key++
	}

	if c == '\n' {
		if t.Next(c) {
			t.Reset()
			return true
		}

		*key = 0
		*index++
	}

	return false
}

func New(input string) *Game {
	game := new(Game)

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	t := transition{}
	state := StateSeeds
	seedCounter := -1
	rangeKey := 0
	rangeIndex := -1

	for scanner.Scan() {
		c := rune(scanner.Bytes()[0])
		switch state {
		case StateSeeds:
			if unicode.IsDigit(c) {
				game.SetSeeds(c, seedCounter)
			}

			if c == ' ' {
				seedCounter++
			}

			if c == '\n' {
				if t.Next(c) {
					t.Reset()
					state = StateSeedToSoil
					continue
				}
			}

			t.Append(c)
		case StateSeedToSoil:
			if parseRange(c, &game.SeedToSoil, &rangeIndex, &rangeKey, &t) {
				rangeIndex = -1
				t.Reset()
				state = StateSoilToFertilizer
				continue
			}

			t.Append(c)
		case StateSoilToFertilizer:
			if parseRange(c, &game.SoilToFertilizer, &rangeIndex, &rangeKey, &t) {
				rangeIndex = -1
				t.Reset()
				state = StateFertilizerToWater
				continue
			}

			t.Append(c)
		case StateFertilizerToWater:
			if parseRange(c, &game.FertilizerToWater, &rangeIndex, &rangeKey, &t) {
				rangeIndex = -1
				t.Reset()
				state = StateWaterToLight
				continue
			}

			t.Append(c)
		case StateWaterToLight:
			if parseRange(c, &game.WaterToLight, &rangeIndex, &rangeKey, &t) {
				rangeIndex = -1
				t.Reset()
				state = StateLightToTemperature
				continue
			}

			t.Append(c)
		case StateLightToTemperature:
			if parseRange(c, &game.LightToTemperature, &rangeIndex, &rangeKey, &t) {
				rangeIndex = -1
				t.Reset()
				state = StateTemperatureToHumidity
				continue
			}

			t.Append(c)
		case StateTemperatureToHumidity:
			if parseRange(c, &game.TemperatureToHumidity, &rangeIndex, &rangeKey, &t) {
				rangeIndex = -1
				t.Reset()
				state = StateHumidityToLocation
				continue
			}

			t.Append(c)
		case StateHumidityToLocation:
			if parseRange(c, &game.HumidityToLocation, &rangeIndex, &rangeKey, &t) {
				rangeIndex = -1
				t.Reset()
			}

			t.Append(c)
		}
	}

	return game
}

package day5

import (
	"math"
)

func Puzzle1(input string) int {
	game := New(input)
	answer := math.MaxInt32
	for _, seed := range game.Seeds {
		target := seed
		for _, _range := range game.SeedToSoil {
			t, found := _range.FindNextMapping(target)
			if found {
				target = t
				break
			}
		}

		for _, _range := range game.SoilToFertilizer {
			t, found := _range.FindNextMapping(target)
			if found {
				target = t
				break
			}
		}

		for _, _range := range game.FertilizerToWater {
			t, found := _range.FindNextMapping(target)
			if found {
				target = t
				break
			}
		}

		for _, _range := range game.WaterToLight {
			t, found := _range.FindNextMapping(target)
			if found {
				target = t
				break
			}
		}

		for _, _range := range game.LightToTemperature {
			t, found := _range.FindNextMapping(target)
			if found {
				target = t
				break
			}
		}

		for _, _range := range game.TemperatureToHumidity {
			t, found := _range.FindNextMapping(target)
			if found {
				target = t
				break
			}
		}

		for _, _range := range game.HumidityToLocation {
			t, found := _range.FindNextMapping(target)
			if found {
				target = t
				break
			}
		}

		if target < answer {
			answer = target
		}
	}

	return answer
}

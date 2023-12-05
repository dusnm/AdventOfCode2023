package day5

import (
	"math"
	"slices"
	"sync"
)

func Puzzle2(input string) int {
	game := New(input)
	answers := make([]int, 0, len(game.Seeds)/2)

	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}

	wg.Add(len(game.Seeds) / 2)
	for i := 0; i < len(game.Seeds); i += 2 {
		seedRange := game.Seeds[i]
		length := game.Seeds[i+1]

		// Brute force Go brrrrr
		go func(
			wg *sync.WaitGroup,
			mu *sync.RWMutex,
			seedRange int,
			length int,
			seedToSoil []Range,
			soilToFertilizer []Range,
			fertilizerToWater []Range,
			waterToLight []Range,
			lightToTemperature []Range,
			temperatureToHumidity []Range,
			humidityToLocation []Range,
		) {
			defer wg.Done()

			answer := math.MaxInt32
			for j := 0; j < length; j++ {
				target := seedRange + j
				for _, _range := range seedToSoil {
					t, found := _range.FindNextMapping(target)
					if found {
						target = t
						break
					}
				}

				for _, _range := range soilToFertilizer {
					t, found := _range.FindNextMapping(target)
					if found {
						target = t
						break
					}
				}

				for _, _range := range fertilizerToWater {
					t, found := _range.FindNextMapping(target)
					if found {
						target = t
						break
					}
				}

				for _, _range := range waterToLight {
					t, found := _range.FindNextMapping(target)
					if found {
						target = t
						break
					}
				}

				for _, _range := range lightToTemperature {
					t, found := _range.FindNextMapping(target)
					if found {
						target = t
						break
					}
				}

				for _, _range := range temperatureToHumidity {
					t, found := _range.FindNextMapping(target)
					if found {
						target = t
						break
					}
				}

				for _, _range := range humidityToLocation {
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

			mu.Lock()
			answers = append(answers, answer)
			mu.Unlock()

			answer = math.MaxInt32
		}(
			wg,
			mu,
			seedRange,
			length,
			game.SeedToSoil,
			game.SoilToFertilizer,
			game.FertilizerToWater,
			game.WaterToLight,
			game.LightToTemperature,
			game.TemperatureToHumidity,
			game.HumidityToLocation,
		)
	}

	wg.Wait()

	return slices.Min(answers)
}

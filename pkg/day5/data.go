package day5

const (
	RangeDestination int = iota
	RangeSource
	RangeLength
)

type (
	Range struct {
		Destination int
		Source      int
		Length      int
	}

	Game struct {
		Seeds                 []int
		SeedToSoil            []Range
		SoilToFertilizer      []Range
		FertilizerToWater     []Range
		WaterToLight          []Range
		LightToTemperature    []Range
		TemperatureToHumidity []Range
		HumidityToLocation    []Range
	}
)

func (r *Range) FindNextMapping(target int) (int, bool) {
	if target < r.Source || target >= r.Source+r.Length {
		return 0, false
	}

	return (r.Destination + r.Length) - (r.Source + r.Length - target), true
}

func (g *Game) SetSeeds(c rune, index int) {
	if len(g.Seeds) <= index {
		g.Seeds = append(g.Seeds, int(c-'0'))
	} else {
		g.Seeds[index] = g.Seeds[index]*10 + int(c-'0')
	}
}

package day6

type (
	Race struct {
		Time     int
		Distance int
	}
)

func (r *Race) IsEnoughToBeatRecord(holdLength int) bool {
	if holdLength >= r.Time || holdLength == 0 {
		return false
	}

	return holdLength*(r.Time-holdLength) > r.Distance
}

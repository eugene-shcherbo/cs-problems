package arrays

func HighestAltitude(gains []int) int {
	highest := 0
	curr := 0

	for _, gain := range gains {
		curr += gain
		highest = max(highest, curr)
	}

	return highest
}

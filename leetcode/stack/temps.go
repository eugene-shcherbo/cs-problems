package stack

func DailyTemperaturesMonoStack(temperatures []int) []int {
	warmerDays := make([]int, len(temperatures))
	tempDecrease := []int{}

	for i, t := range temperatures {
		for len(tempDecrease) > 0 && temperatures[tempDecrease[len(tempDecrease)-1]] < t {
			coldIdx := tempDecrease[len(tempDecrease)-1]
			warmerDays[coldIdx] = i - coldIdx
			tempDecrease = tempDecrease[:len(tempDecrease)-1]
		}
		tempDecrease = append(tempDecrease, i)
	}

	return warmerDays
}

func DailyTemperaturesJumps(temperatures []int) []int {
	warmerDays := make([]int, len(temperatures))

	for i := len(temperatures) - 2; i >= 0; i-- {
		j := i + 1
		for temperatures[i] >= temperatures[j] && warmerDays[j] != 0 {
			j += warmerDays[j]
		}

		if warmerDays[j] == 0 && temperatures[i] >= temperatures[j] {
			warmerDays[i] = 0
		} else {
			warmerDays[i] = j - i
		}
	}

	return warmerDays
}

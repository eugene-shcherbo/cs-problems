package stack

func RemoveKdigits(num string, k int) string {
	newNum := make([]rune, 0)

	for _, d := range num {
		for len(newNum) > 0 && k > 0 && newNum[len(newNum)-1] > d {
			k--
			newNum = newNum[:len(newNum)-1]
		}

		if len(newNum) > 0 || d != '0' {
			newNum = append(newNum, d)
		}
	}

	for len(newNum) > 0 && k > 0 {
		k--
		newNum = newNum[:len(newNum)-1]
	}

	if len(newNum) > 0 {
		return string(newNum)
	}

	return "0"
}

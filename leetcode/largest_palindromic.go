package leetcode

func LargestPalindromic(num string) string {
	counts := map[rune]int{}

	for _, r := range num {
		counts[r]++
	}

	start := make([]rune, 0)
	end := make([]rune, 0)

	for d := '9'; d >= '0'; d-- {
		if d != '0' || len(start) > 0 {
			for i := 0; i < counts[d]/2; i++ {
				start = append(start, d)
				end = append(end, d)
			}
		}
	}

	for d := '9'; d >= '0'; d-- {
		if counts[d]%2 != 0 {
			end = append(end, d)
			break
		}
	}

	res := make([]rune, len(start))
	copy(res, start)

	for i := len(end) - 1; i >= 0; i-- {
		res = append(res, end[i])
	}

	if len(res) == 0 && counts['0'] > 0 {
		return "0"
	}

	return string(res)
}

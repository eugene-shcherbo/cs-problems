package leetcode

func MinAddToMakeValid(s string) int {
	closing := 0
	adds := 0

	for _, r := range s {
		if r == '(' {
			closing++
		} else if closing == 0 {
			adds++
		} else {
			closing--
		}
	}

	return adds + closing
}

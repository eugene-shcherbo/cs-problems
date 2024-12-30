package leetcode

func CheckInclusion(s1 string, s2 string) bool {
	var s1Counts = map[byte]int{}
	for _, r := range s1 {
		s1Counts[byte(r)]++
	}

	var s2Counts = map[byte]int{}
	matched := 0
	for l, r := 0, 0; r < len(s2); r++ {
		s2Counts[s2[r]]++

		if s1Counts[s2[r]] == s2Counts[s2[r]] {
			matched++
		}

		if r >= len(s1) {
			s2Counts[s2[l]]--
			if s2Counts[s2[l]] == s1Counts[s2[l]]-1 {
				matched--
			}
			l++
		}

		if matched == len(s1Counts) {
			return true
		}
	}

	return false
}

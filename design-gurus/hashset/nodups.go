package hashset

func LongestSubstringWithoutDups(s string) int {
	var duplicates = map[byte]bool{}

	maxLength := 0
	for l, r := 0, 0; r < len(s); r++ {
		if !duplicates[s[r]] {
			newLength := r - l + 1
			if newLength > maxLength {
				maxLength = newLength
			}
		} else {
			for duplicates[s[r]] {
				duplicates[s[l]] = false
				l++
			}
		}

		duplicates[s[r]] = true
	}

	return maxLength
}

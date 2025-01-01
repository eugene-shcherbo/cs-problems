package slidingwindow

func CharacterReplacement(s string, k int) int {
	bestSize := 0
	counts := [26]int{}
	maxCount := 0

	for l, r := 0, 0; r < len(s); r++ {
		counts[s[r]-'A']++
		maxCount = max(maxCount, counts[s[r]-'A'])

		if r-l+1-maxCount > k {
			counts[s[l]-'A']--
			l++
		}

		bestSize = r - l + 1
	}

	return bestSize
}

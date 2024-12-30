package leetcode

func FindSubstring(s string, words []string) []int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	indicies := make([]int, 0)
	wordSize := len(words[0])
	winSize := len(words) * wordSize
	matched := 0
	required := len(counts)

	for l, r := 0, 0; r <= len(s)-wordSize; r += wordSize {
		endWord := s[r : r+wordSize]
		counts[endWord]--

		if counts[endWord] == 0 {
			matched++
		}

		if r >= winSize {
			firstWord := s[l : l+wordSize]
			counts[firstWord]++
			if counts[firstWord] == 1 {
				matched--
			}
			l += wordSize
		}

		if matched == required {
			indicies = append(indicies, l)
		}
	}

	return indicies
}

func FindSubstringSlidingWindow(s string, words []string) []int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	indicies := make([]int, 0)
	wordSize := len(words[0])
	winSize := len(words) * wordSize

	for i := 0; i < wordSize; i++ {
		sCounts := make(map[string]int)
		matched := 0

		for l, r := i, i; r <= len(s)-wordSize; r += wordSize {
			endWord := s[r : r+wordSize]
			sCounts[endWord]++

			if sCounts[endWord] == counts[endWord] {
				matched++
			}

			if r-l+1 > winSize {
				firstWord := s[l : l+wordSize]
				sCounts[firstWord]--

				if sCounts[firstWord] == counts[firstWord]-1 {
					matched--
				}

				l += wordSize
			}

			if matched == len(counts) {
				indicies = append(indicies, l)
			}
		}
	}

	return indicies
}

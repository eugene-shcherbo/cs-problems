package leetcode

func BackspaceCompare(s string, t string) bool {
	for l, r := len(s)-1, len(t)-1; l >= 0 || r >= 0; l, r = l-1, r-1 {
		if l < 0 || r < 0 {
			return false
		}

		l = getNextIndex(s, l)
		r = getNextIndex(t, r)

		if l < 0 && r < 0 {
			return true
		}

		if l < 0 || r < 0 || s[l] != t[r] {
			return false
		}
	}

	return true
}

func getNextIndex(s string, idx int) int {
	toRemove := 0
	for idx >= 0 && (s[idx] == '#' || toRemove > 0) {
		if s[idx] == '#' {
			toRemove++
		} else {
			toRemove--
		}
		idx--
	}
	return idx
}

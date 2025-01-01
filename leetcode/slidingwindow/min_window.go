package slidingwindow

type Range struct {
	s, e int
}

func (r *Range) Len() int {
	return r.e - r.s + 1
}

func MinWindow(s string, t string) string {
	counts := make(map[byte]int)
	for _, r := range t {
		counts[byte(r)]++
	}

	matched := 0
	minRange := Range{0, len(s)}
	for l, r := 0, 0; r < len(s); r++ {
		if v, ok := counts[s[r]]; ok {
			counts[s[r]]--
			if v == 1 {
				matched++
			}
		}

		for matched == len(counts) {
			if r-l+1 < minRange.Len() {
				minRange.s, minRange.e = l, r
			}

			if v, ok := counts[s[l]]; ok {
				if v == 0 {
					matched--
				}
				counts[s[l]]++
			}

			l++
		}
	}

	if minRange.Len() > len(s) {
		return ""
	}

	return s[minRange.s : minRange.e+1]
}

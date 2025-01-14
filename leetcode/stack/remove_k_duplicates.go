package stack

import "strings"

type letterInfo struct {
	letter rune
	count  int
}

func RemoveKDuplicates(s string, k int) string {
	keptLetters := make([]*letterInfo, 0)

	for _, r := range s {

		if len(keptLetters) > 0 && keptLetters[len(keptLetters)-1].letter == r {
			keptLetters[len(keptLetters)-1].count++
		} else {
			keptLetters = append(keptLetters, &letterInfo{r, 1})
		}

		if keptLetters[len(keptLetters)-1].count == k {
			keptLetters = keptLetters[:len(keptLetters)-1]
		}
	}

	var newString strings.Builder
	for _, i := range keptLetters {
		for j := 0; j < i.count; j++ {
			newString.WriteRune(i.letter)
		}
	}

	return newString.String()
}

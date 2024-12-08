package leetcode

import (
	"sort"
	"unicode"
)

func SortVowels(s string) string {
	vowels := make([]rune, 0)

	for _, r := range s {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	res := make([]rune, 0)
	vowelIdx := 0
	for _, r := range s {
		if !isVowel(r) {
			res = append(res, r)
		} else {
			res = append(res, vowels[vowelIdx])
			vowelIdx++
		}
	}

	return string(res)
}

func isVowel(r rune) bool {
	low := unicode.ToLower(r)
	return low == 'a' || low == 'e' || low == 'i' || low == 'o' || low == 'u'
}

package twopointers

import (
	"math"

	"github.com/eugene-shcherbo/cs-problems/utils"
)

func ShortestDistance(wordsDict []string, word1 string, word2 string) int {
	word1Indicies := make([]int, 0)
	word2Indicies := make([]int, 0)

	for i, w := range wordsDict {
		if w == word1 {
			word1Indicies = append(word1Indicies, i)
		}
		if w == word2 {
			word2Indicies = append(word2Indicies, i)
		}
	}

	minDist := math.MaxInt
	for idx1, idx2 := 0, 0; idx1 < len(word1Indicies) && idx2 < len(word2Indicies); {
		minDist = utils.Min(minDist, utils.AbsInt(word1Indicies[idx1]-word2Indicies[idx2]))
		if word1Indicies[idx1] < word2Indicies[idx2] {
			idx1++
		} else {
			idx2++
		}
	}

	return minDist
}

package intervals

import "github.com/eugene-shcherbo/cs-problems/utils"

func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	intersections := make([][]int, 0)

	for l, r := 0, 0; l < len(firstList) && r < len(secondList); {
		first := firstList[l]
		second := secondList[r]
		if secondList[r][0] < first[0] {
			first, second = second, first
		}

		if first[1] >= second[0] {
			intersection := []int{utils.Max(firstList[l][0], secondList[r][0]), utils.Min(firstList[l][1], secondList[r][1])}
			intersections = append(intersections, intersection)
		}

		if firstList[l][1] >= secondList[r][1] {
			r++
		} else {
			l++
		}
	}

	return intersections
}

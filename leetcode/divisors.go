package leetcode

import "github.com/eugene-shcherbo/cs-problems/utils"

func MinimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	total := uniqueCnt1 + uniqueCnt2
	lcm := divisor1 * divisor2 / utils.Gcd(divisor1, divisor2)

	right := 1
	for {
		v1, v2 := countSets(divisor1, divisor2, right)
		if canFill(right, lcm, v1, v2, uniqueCnt1, uniqueCnt2) {
			break
		}
		right *= 2
	}

	left := total

	for left <= right {
		mid := (left + right) / 2

		v1, v2 := countSets(divisor1, divisor2, mid)

		if canFill(mid, lcm, v1, v2, uniqueCnt1, uniqueCnt2) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func countSets(div1, div2, maxValue int) (int, int) {
	first := maxValue - (maxValue / div1)
	second := maxValue - (maxValue / div2)

	return first, second
}

func canFill(val, commonDivisor, v1, v2, uniq1, uniq2 int) bool {
	return val-(val/commonDivisor) >= uniq1+uniq2 && v1 >= uniq1 && v2 >= uniq2
}

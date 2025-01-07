package intervals

import (
	"math"

	"github.com/eugene-shcherbo/cs-problems/leetcode"
	"github.com/eugene-shcherbo/cs-problems/utils"
)

func EmployeeFreeTime(schedule [][]*leetcode.Interval) []*leetcode.Interval {
	freeIntervals := make([][]*leetcode.Interval, len(schedule))

	for i, emp := range schedule {
		freeIntervals[i] = make([]*leetcode.Interval, 0)
		freeIntervals[i] = append(freeIntervals[i], &leetcode.Interval{Start: math.MinInt, End: emp[0].Start})

		for j := 1; j < len(emp); j++ {
			interval := &leetcode.Interval{Start: emp[j-1].End, End: emp[j].Start}
			freeIntervals[i] = append(freeIntervals[i], interval)
		}

		freeIntervals[i] = append(freeIntervals[i], &leetcode.Interval{Start: emp[len(emp)-1].End, End: math.MaxInt})
	}

	intersections := freeIntervals[0]
	for i := 1; i < len(freeIntervals); i++ {
		first, second := 0, 0
		newIntersections := make([]*leetcode.Interval, 0)
		for first < len(intersections) && second < len(freeIntervals[i]) {
			start := utils.Max(intersections[first].Start, freeIntervals[i][second].Start)
			end := utils.Min(intersections[first].End, freeIntervals[i][second].End)

			if start < end {
				newIntersections = append(newIntersections, &leetcode.Interval{Start: start, End: end})
			}

			if intersections[first].End < freeIntervals[i][second].End {
				first++
			} else {
				second++
			}
		}
		intersections = newIntersections
	}

	if len(intersections) > 0 {
		if intersections[0].Start == math.MinInt {
			intersections = intersections[1:]
		}
	}

	if len(intersections) > 0 {
		if intersections[len(intersections)-1].End == math.MaxInt {
			intersections = intersections[:len(intersections)-1]
		}
	}

	return intersections
}

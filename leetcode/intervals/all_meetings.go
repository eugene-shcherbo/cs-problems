package intervals

import "sort"

func CanAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	for i, interval := range intervals[:len(intervals)-1] {
		if interval[1] > intervals[i+1][0] {
			return false
		}
	}

	return true
}

package slidingwindow

func SubarraysWithKDistinct(nums []int, k int) int {
	countsTotal := make(map[int]int)
	countsMid := make(map[int]int)

	numOfSubArrays := 0

	for l, m, r := 0, 0, 0; r < len(nums); r++ {
		countsTotal[nums[r]]++
		countsMid[nums[r]]++

		for len(countsMid) == k {
			countsMid[nums[m]]--
			if countsMid[nums[m]] == 0 {
				delete(countsMid, nums[m])
			}
			m++
		}

		for len(countsTotal) > k {
			countsTotal[nums[l]]--
			if countsTotal[nums[l]] == 0 {
				delete(countsTotal, nums[l])
			}
			l++
		}

		if len(countsTotal) == k {
			numOfSubArrays += m - l
		}
	}

	return numOfSubArrays
}

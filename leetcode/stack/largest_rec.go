package stack

// Observations:
// 1. If we found smaller height, we don't care in higher heights starting from this point (they can be removed) - mono stack
// 2. Increasing towers are stored together (in mono stack), so we can count running area while deleting them
// 3. Next: 2,1,1,1,2,3
func LargestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	maxArea := 0
	idxOfPeaks := make([]int, 0)

	for i, h := range heights {
		for len(idxOfPeaks) > 0 && heights[idxOfPeaks[len(idxOfPeaks)-1]] > h {
			peakHight := heights[idxOfPeaks[len(idxOfPeaks)-1]]
			idxOfPeaks = idxOfPeaks[:len(idxOfPeaks)-1]

			right := i - 1
			left := 0
			if len(idxOfPeaks) > 0 {
				left = idxOfPeaks[len(idxOfPeaks)-1] + 1
			}

			maxArea = max(maxArea, (right-left+1)*peakHight)
		}

		idxOfPeaks = append(idxOfPeaks, i)
	}

	return maxArea
}

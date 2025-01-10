package cyclicsort

func FindDuplicateCyclicSort(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i]-1 != i && nums[i] == nums[nums[i]-1] {
			return nums[i]
		} else if nums[i]-1 != i {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	return -1
}

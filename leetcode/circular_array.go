package leetcode

func CircularArrayLoopSet(nums []int) bool {
	n := len(nums)
	visited := make([]bool, n)

	for i := range nums {
		if visited[i] {
			continue
		}

		localVisited := make([]bool, n)

		j := i
		isForward := nums[j] > 0
		for ; isForward == (nums[j] > 0) && !localVisited[j]; j = calcIndex(n, j, nums[j]) {
			visited[j] = true
			localVisited[j] = true
		}

		if nums[j]%n != 0 && -nums[j]%n != 0 && isForward == (nums[j] > 0) {
			return true
		}
	}

	return false
}

func CircularArrayLoopSlowFast(nums []int) bool {
	n := len(nums)
	visited := make([]bool, n)

	for i, v := range nums {
		if visited[i] {
			continue
		}

		isForward := nums[i] > 0
		slow := i
		fast := calcIndex(len(nums), i, v)

		for isForward == (nums[fast] > 0) && slow != fast {
			visited[slow] = true

			if isForward != (nums[fast] > 0) || isForward != (nums[calcIndex(n, fast, nums[fast])] > 0) {
				break
			}

			slow = calcIndex(len(nums), slow, nums[slow])
			fast = calcIndex(len(nums), fast, nums[fast])
			fast = calcIndex(len(nums), fast, nums[fast])
		}

		if nums[fast]%n != 0 && -nums[fast]%n != 0 && slow == fast {
			return true
		}
	}

	return false
}

func calcIndex(size int, curr int, shift int) int {
	if shift < 0 {
		shift = -(-shift % size) + size
	}

	return (curr + shift) % size
}

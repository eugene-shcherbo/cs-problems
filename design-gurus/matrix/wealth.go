package matrix

func MaxWealth(accounts [][]int) int {
	maxWealth := 0

	for i := 0; i < len(accounts); i++ {
		wealth := 0
		for _, money := range accounts[i] {
			wealth += money
		}
		maxWealth = max(maxWealth, wealth)
	}

	return maxWealth
}

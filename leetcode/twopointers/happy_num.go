package twopointers

func IsHappy(n int) bool {
	slow := n
	fast := calcSquareSum(n)

	for fast != 1 && slow != fast {
		slow = calcSquareSum(slow)
		fast = calcSquareSum(calcSquareSum(fast))
	}

	return fast == 1
}

func calcSquareSum(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

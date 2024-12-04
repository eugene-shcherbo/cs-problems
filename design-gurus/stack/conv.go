package stack

import (
	"fmt"
	"strings"
)

func DecimalToBinary(num int) string {
	var binaryDigits []int

	for num != 0 {
		binaryDigits = append(binaryDigits, num%2)
		num /= 2
	}

	var res strings.Builder
	for i := len(binaryDigits) - 1; i >= 0; i-- {
		res.WriteString(fmt.Sprintf("%d", binaryDigits[i]))
	}

	return res.String()
}

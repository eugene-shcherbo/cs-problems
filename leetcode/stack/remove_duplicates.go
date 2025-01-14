package stack

func RemoveDuplicates(s string, k int) string {
	keptChars := []rune{}
	for _, v := range s {
		if len(keptChars) > 0 && keptChars[len(keptChars)-1] == v {
			keptChars = keptChars[:len(keptChars)-1]
		} else {
			keptChars = append(keptChars, v)
		}
	}

	return string(keptChars)
}

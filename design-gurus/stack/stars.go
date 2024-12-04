package stack

func RemoveStars(s string) string {
	res := make([]rune, 0)

	for _, r := range s {
		if r != '*' {
			res = append(res, r)
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}

	return string(res)
}

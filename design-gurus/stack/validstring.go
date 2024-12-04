package stack

func RemoveDuplicates(s string) string {
	var res []byte

	for i := 0; i < len(s); {
		if i != len(s)-1 && s[i] == s[i+1] {
			i += 2
			continue
		}

		if len(res) > 0 && res[len(res)-1] == s[i] {
			res = res[:len(res)-1]
			i++
			continue
		}

		res = append(res, s[i])
		i++
	}

	return string(res)
}

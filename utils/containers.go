package utils

func Reverse[T any](data []T) []T {
	for l, r := 0, len(data)-1; l < r; l, r = l+1, r-1 {
		data[l], data[r] = data[r], data[l]
	}

	return data
}

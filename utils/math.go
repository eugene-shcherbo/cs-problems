package utils

import "cmp"

func AbsInt[T int](a T) T {
	if a < 0 {
		return -a
	}

	return a
}

func Gcd[T int](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Min[T cmp.Ordered](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Max[T cmp.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

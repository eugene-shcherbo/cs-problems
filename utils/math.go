package utils

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

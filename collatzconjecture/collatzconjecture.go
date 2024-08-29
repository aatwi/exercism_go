package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Invalid Number")
	}
	return collatzConjecture(n), nil
}

func collatzConjecture(n int) int {
	if n == 1 {
		return 0
	}
	if isEven(n) {
		return 1 + collatzConjecture(n/2)
	}
	return 1 + collatzConjecture(3*n+1)
}

func isEven(n int) bool {
	return n%2 == 0
}

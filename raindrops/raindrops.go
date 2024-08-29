package raindrops

import "strconv"

func Convert(number int) string {
	result := getDivisibleByText(number, "Pling", divisibleBy(3))
	result += getDivisibleByText(number, "Plang", divisibleBy(5))
	result += getDivisibleByText(number, "Plong", divisibleBy(7))

	if len(result) == 0 {
		return strconv.Itoa(number)
	}
	return result
}

func divisibleBy(by int) func(int) bool {
	return func(number int) bool {
		return number%by == 0
	}
}

func getDivisibleByText(num int, text string, predicate func(int) bool) string {
	if predicate(num) {
		return text
	}
	return ""
}

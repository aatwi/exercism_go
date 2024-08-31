package strain

func Keep[T any](cp []T, predicate func(T) bool) []T {
	result := []T{}
	for _, element := range cp {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}

func Discard[T any](cp []T, predicate func(T) bool) []T {
	result := []T{}
	for _, element := range cp {
		if predicate(element) == false {
			result = append(result, element)
		}
	}
	return result
}

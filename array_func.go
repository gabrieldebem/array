package arrays

func Map[T, U any](slice []T, f func(T) U) []U {
	var result []U
	for _, value := range slice {
		result = append(result, f(value))
	}
	return result
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func Reduce[T any](slice []T, f func(T, T) T) T {
	result := slice[0]
	for _, value := range slice[1:] {
		result = f(result, value)
	}
	return result
}

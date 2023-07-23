package array

// Map applies a function to each element of a slice and returns a new slice
func Map[T, U any](slice []T, f func(T) U) []U {
	var result []U
	for _, value := range slice {
		result = append(result, f(value))
	}
	return result
}

// Filter returns a new slice containing only the elements of the slice that satisfy the predicate
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

// Reduce applies a function to each element of a slice and returns a new slice
func Reduce[T any](slice []T, f func(T, T) T) T {
	result := slice[0]
	for _, value := range slice[1:] {
		result = f(result, value)
	}
	return result
}

package gofun

// Any checks if any element in the slice satisfies the provided predicate function.
// It returns true if at least one element meets the condition, otherwise false.
func Any[T any](ts []T, predicate func(T) bool) bool {
	for _, t := range ts {
		if predicate(t) {
			return true
		}
	}
	return false
}

package main

// None checks if no elements in the slice satisfy the given predicate function.
// Returns true if none matches, otherwise false.
func None[T any](ts []T, predicate func(T) bool) bool {
	for _, t := range ts {
		if predicate(t) {
			return false
		}
	}
	return true
}

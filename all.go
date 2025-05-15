package main

// All checks if all elements in the slice satisfy the given predicate function.
// Returns true if all pass, otherwise false.
func All[T any](ts []T, predicate func(T) bool) bool {
	for _, t := range ts {
		if !predicate(t) {
			return false
		}
	}
	return true
}

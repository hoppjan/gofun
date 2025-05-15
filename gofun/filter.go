package gofun

// Filter filters a slice for the given predicate and returns a new slice
// that only contains the ones that match the predicate.
func Filter[T any](ts []T, predicate func(T) bool) []T {
	var us = make([]T, 0, len(ts))
	for _, t := range ts {
		if predicate(t) {
			us = append(us, t)
		}
	}
	return us
}

package main

// Find returns the first t in slice ts that equals find, or nil if no such element exists.
func Find[T comparable](ts []T, find T) *T {
	for _, t := range ts {
		if t == find {
			return &t
		}
	}
	return nil
}

// FindMatch returns the first t in slice ts that the matcher identifies to be matching (returning true).
// If no matching element is found, nil is returned.
func FindMatch[T any](ts []T, matcher func(t T) bool) *T {
	for _, t := range ts {
		if matcher(t) {
			return &t
		}
	}
	return nil
}

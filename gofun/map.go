package gofun

// Map applies the function f on each element t in ts, transforming each t
// from type T to type U, returning a new list of Us.
func Map[T any, U any](ts []T, f func(t T) U) []U {
	var us = make([]U, 0, len(ts))
	for _, t := range ts {
		us = append(us, f(t))
	}
	return us
}

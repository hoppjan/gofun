package gofun

// Reduce applies a reduction function f to each element in the slice ts, combining them into a single result.
func Reduce[T any, U any](ts []T, initial U, f func(t T, acc U) U) U {
	u := initial
	for _, t := range ts {
		u = f(t, u)
	}
	return u
}

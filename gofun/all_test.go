package gofun

import "testing"

func TestAll(t *testing.T) {
	type args[T any] struct {
		ts        []T
		predicate func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			"All should return true, given the predicate matches all",
			args[int]{[]int{2, 4, 6, 8}, func(t int) bool { return t%2 == 0 }},
			true,
		},
		{
			"All should return false, given the predicate does not match all",
			args[int]{[]int{2, 4, 6, 9}, func(t int) bool { return t%2 == 1 }},
			false,
		},
		{
			"All should return true, given the slice is empty",
			args[int]{[]int{}, func(_ int) bool { return false }},
			true,
		},
		{
			"All should return true, given the slice is nil",
			args[int]{nil, func(_ int) bool { return false }},
			true,
		},
		{
			"All should return true, given the slice and predicate are nil",
			args[int]{nil, nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.ts, tt.args.predicate); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

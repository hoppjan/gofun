package gofun

import "testing"

func TestNone(t *testing.T) {
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
			"None should return false, given the predicate matches",
			args[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(t int) bool { return t%2 == 0 }},
			false,
		},
		{
			"None should return false, given the predicate matches even only one",
			args[int]{[]int{1, 3, 5, 7, 9, 10}, func(t int) bool { return t%2 == 0 }},
			false,
		},
		{
			"None should return true, given the predicate does not match",
			args[int]{[]int{0, 2, 4, 6, 8, 10}, func(t int) bool { return t%2 == 1 }},
			true,
		},
		{
			"None should return true, given the slice is empty",
			args[int]{[]int{}, func(_ int) bool { return true }},
			true,
		},
		{
			"None should return true, given the slice is nil",
			args[int]{nil, func(_ int) bool { return true }},
			true,
		},
		{
			"None should return true, given the slice and predicate are nil",
			args[int]{nil, nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := None(tt.args.ts, tt.args.predicate); got != tt.want {
				t.Errorf("None() = %v, want %v", got, tt.want)
			}
		})
	}
}

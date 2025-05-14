package main

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		ts        []T
		predicate func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			"Filter even numbers should return even numbers",
			args[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(t int) bool { return t%2 == 0 }},
			[]int{2, 4, 6, 8},
		},
		{
			"Filter odd numbers should return odd numbers",
			args[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(t int) bool { return t%2 == 1 }},
			[]int{1, 3, 5, 7, 9},
		},
		{
			"Filter that matches nothing should return empty slice",
			args[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(t int) bool { return t > 10 }},
			[]int{},
		},
		{
			"Filter empty slice should return empty slice",
			args[int]{[]int{}, func(t int) bool { return true }},
			[]int{},
		},
		{
			"Filter nil should return empty slice",
			args[int]{nil, func(t int) bool { return true }},
			[]int{},
		},
		{
			"Filter with nil predicate should not panic if slice is empty",
			args[int]{[]int{}, nil},
			[]int{},
		},
		{
			"Filter with nil predicate should not panic if slice nil",
			args[int]{nil, nil},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.ts, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
	if x := Filter[string](nil, nil); len(x) != 0 {
		t.Errorf("Filter() = %v, want %v", x, []string{})
	}
}

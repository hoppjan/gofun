package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReduce_int_int(t *testing.T) {
	type args[T interface{}, U interface{}] struct {
		ts      []T
		initial U
		f       func(t T, u U) U
	}
	type testCase[T interface{}, U interface{}] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, int]{
		{
			name: "Add up ints",
			args: args[int, int]{
				ts:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				initial: 0,
				f:       func(t int, u int) int { return t + u },
			},
			want: 45,
		},
		{
			name: "Add up empty list to initial value",
			args: args[int, int]{
				ts:      []int{},
				initial: 0,
				f:       func(t int, u int) int { return t + u },
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.ts, tt.args.initial, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce_int_string(t *testing.T) {
	type args[T interface{}, U interface{}] struct {
		ts      []T
		initial U
		f       func(t T, u U) U
	}
	type testCase[T interface{}, U interface{}] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, string]{
		{
			name: "Reduce ints to string",
			args: args[int, string]{
				ts:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				initial: "",
				f:       func(t int, acc string) string { return acc + fmt.Sprint(t) },
			},
			want: "123456789",
		},
		{
			name: "Reduce empty list to initial string",
			args: args[int, string]{
				ts:      []int{},
				initial: "",
				f:       func(t int, acc string) string { return acc + fmt.Sprint(t) },
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.ts, tt.args.initial, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

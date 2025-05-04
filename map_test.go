package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap_int_int(t *testing.T) {
	type args[T interface{}, U interface{}] struct {
		ts []T
		f  func(t T) U
	}
	type testCase[T interface{}, U interface{}] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[int, int]{
		{
			name: "Map ints times ten",
			args: args[int, int]{
				ts: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				f:  func(n int) int { return n * 10 },
			},
			want: []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
		{
			name: "Map empty list returns empty list",
			args: args[int, int]{
				ts: []int{},
				f:  func(t int) int { return t },
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.args.ts, tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_int_string(t *testing.T) {
	type args[T interface{}, U interface{}] struct {
		ts []T
		f  func(t T) U
	}
	type testCase[T interface{}, U interface{}] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[int, string]{
		{
			name: "Map ints to strings",
			args: args[int, string]{
				ts: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				f:  func(n int) string { return fmt.Sprint(n) },
			},
			want: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
		{
			name: "Map empty list returns empty list",
			args: args[int, string]{
				ts: []int{},
				f:  func(t int) string { return string(rune(t)) },
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.args.ts, tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

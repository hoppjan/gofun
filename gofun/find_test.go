package gofun

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type args[T comparable] struct {
		ts   []T
		find T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want *T
	}
	abc := "abc"
	tests := []testCase[string]{
		{"return *abc if abc in slice", args[string]{[]string{"", "abc"}, "abc"}, &abc},
		{"return nil if def not in slice", args[string]{[]string{"", "abc"}, "def"}, nil},
		{"return nil if slice nil", args[string]{nil, "def"}, nil},
		{"return nil if slice empty", args[string]{[]string{}, "abc"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.ts, tt.args.find); &tt.args.find == got {
				t.Errorf("Find() = %v, want %v", got, &tt.args.find)
			}
		})
	}
}

func TestFindMatching(t *testing.T) {
	type args[T any] struct {
		ts      []T
		matcher func(t T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}
	abc := "abc"
	tests := []testCase[string]{
		{
			name: "return *abc if abc in slice",
			args: args[string]{[]string{"", "abcd", "abc"}, func(s string) bool { return s == "abc" }},
			want: &abc,
		},
		{
			name: "return nil if def not in slice",
			args: args[string]{[]string{"", "abc"}, func(t string) bool { return t == "def" }},
			want: nil,
		},
		{
			name: "return nil if slice nil",
			args: args[string]{nil, func(t string) bool { return t == "def" }},
			want: nil,
		},
		{
			name: "return nil if slice empty",
			args: args[string]{[]string{}, func(t string) bool { return t == "abc" }},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMatch(tt.args.ts, tt.args.matcher); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMatch() = %v, want %v", got, tt.want)
			}
			t.Logf("Find() found: %v", tt.want)
		})
	}
}

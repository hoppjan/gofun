package main

import (
	"fmt"
	"github.com/hoppjan/gofun/gofun"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers = gofun.Map(numbers, func(n int) int { return n * 10 })
	fmt.Println(numbers)
	filtered := gofun.Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println(filtered)
	reduced := gofun.Reduce(filtered, 0, func(n int, acc int) int { return n + acc })
	fmt.Println(reduced)
	fmt.Println(*gofun.Find(numbers, 10))
	fmt.Println(*gofun.FindMatch(numbers, func(n int) bool { return n%2 == 0 }))
	fmt.Println(gofun.Any(numbers, func(n int) bool { return n == 5 }))
	fmt.Println(gofun.All(numbers, func(n int) bool { return n%2 == 0 }))
	fmt.Println(gofun.None(numbers, func(n int) bool { return n%2 == 0 }))
}

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers = Map(numbers, func(n int) int { return n * 10 })
	fmt.Println(numbers)
	filtered := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println(filtered)
	reduced := Reduce(filtered, 0, func(n int, acc int) int { return n + acc })
	fmt.Println(reduced)
	fmt.Println(*Find(numbers, 10))
	fmt.Println(*FindMatch(numbers, func(n int) bool { return n%2 == 0 }))
	fmt.Println(Any(numbers, func(n int) bool { return n == 5 }))
	fmt.Println(All(numbers, func(n int) bool { return n%2 == 0 }))
	fmt.Println(None(numbers, func(n int) bool { return n%2 == 0 }))
}

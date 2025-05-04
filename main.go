package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers = Map(numbers, func(n int) int { return n * 10 })
	fmt.Println(numbers)
	reduced := Reduce(numbers, 0, func(n int, acc int) int { return n + acc })
	fmt.Println(reduced)
}

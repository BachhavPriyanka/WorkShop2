/*Q4) Write a function that takes a slice of integers and a function and returns the smallest and largest integers in the slice that satisfy the function.
input:	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
output:
Min: 2, Max: 10
Min: 1, Max: 9*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min, max := minMax(numbers, isEven)
	fmt.Printf("Min: %d, Max: %d\n", min, max) // Output: Min: 2, Max: 10

	min, max = minMax(numbers, isOdd)
	fmt.Printf("Min: %d, Max: %d\n", min, max) // Output: Min: 1, Max: 9
}

func minMax(numbers []int, predicate func(int) bool) (int, int) {
	min := 0
	max := 0
	first := true
	for _, n := range numbers {
		if predicate(n) {
			if first {
				min = n
				max = n
				first = false
			} else {
				if n < min {
					min = n
				}
				if n > max {
					max = n
				}
			}
		}
	}
	return min, max
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 == 1
}

/*
Write a function Reverse that reverses the order of elements in a slice. The function
should be generic and work for any type.
*/
package main

import "fmt"

func reverse[T any](numbers []T) {
	start := 0
	end := len(numbers) - 1

	for start < end {
		numbers[start], numbers[end] = numbers[end], numbers[start]
		start = start + 1
		end = end - 1
	}
}

func main() {
	sliceOfNumbers := []int{5, 10, 15, 20, 25}
	fmt.Println("Slice of Numbers:", sliceOfNumbers)
	reverse(sliceOfNumbers)
	fmt.Println("Reverse Order :", sliceOfNumbers)

	sliceOfString := []string{"A", "B", "C"}
	reverse(sliceOfString)
	fmt.Println(sliceOfString)
}

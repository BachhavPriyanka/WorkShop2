/*Q3)Write a function called "filter" that takes in a slice of integers and a function called "predicate".
The function should return a new slice containing only the elements from the original slice that satisfy the predicate function, in the same order as the original slice.
input:	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
Output:
[2 4 6 8 10]
[1 3 5 7 9]*/

package main

import "fmt"

func filter(nums []int, predicate func(int) bool) []int {
	newNumberSlice := []int{}
	for _, val := range nums {
		if predicate(val) {
			newNumberSlice = append(newNumberSlice, val)
		}
	}
	return newNumberSlice
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := filter(numbers, func(num int) bool {
		return num%2 == 0
	})
	slice2 := filter(numbers, func(num int) bool {
		return num%2 != 0
	})
	fmt.Println(slice1)
	fmt.Println(slice2)

}

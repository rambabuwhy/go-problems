/*

  Write a function that takes in two non-empty arrays of integers, finds the
  pair of numbers (one from each array) whose absolute difference is closest to
  zero, and returns an array containing these two numbers, with the number from
  the first array in the first position.
  
  */


package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int{-1, 5, 10, 20, 28, 3}
	b := []int{26, 134, 135, 15, 17}
	fmt.Println(SmallestDifference(a, b))
}
func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	fmt.Println(array1, array2)

	index1 := 0
	index2 := 0

	result := []int{}
	small, current := math.MaxInt32, math.MaxInt32

	for index1 < len(array1) && index2 < len(array2) {
		First, Second := array1[index1], array2[index2]

		if First < Second {
			current = Second - First
			index1 = index1 + 1
		} else if First > Second {
			current = First - Second
			index2 = index2 + 1
		} else {
			return []int{First, Second}
		}

		if small > current {
			small = current
			result = []int{First, Second}
		}
	}
	return result
}

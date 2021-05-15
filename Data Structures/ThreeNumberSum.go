package main

import "sort"
func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	result := [][]int{}
	sort.Ints(array)
	for i:=0; i<len(array)-2; i++{
		left := i+1
		right := len(array)-1
		for left < right{
			currSum := array[i] + array[left] + array[right]
			if currSum == target{
				result = append(result, []int{array[i], array[left], array[right]})
				left = left + 1
				right = right - 1
				
			} else if currSum < target {
				left = left +1
			} else if currSum > target {
				right = right -1
			}
		}
	}
	return result
}

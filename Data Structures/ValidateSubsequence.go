/*

  Given two non-empty arrays of integers, write a function that determines
  whether the second array is a subsequence of the first one.
  
*/

package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	
	index1 := 0;
	index2 := 0;
	
	for index1 < len(array) && index2 < len(sequence){
		if array[index1] == sequence[index2] {
			index2 = index2 + 1
		} 
		
		index1 = index1 + 1
	}
	return index2 == len(sequence)
}

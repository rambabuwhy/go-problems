package main

func GetPermutations(array []int) [][]int {
	// Write your code here.
	result := [][]int{}
	util(array, 0, &result)
	return result
}

func util( array []int, index int, result *[][]int){
	if index == len(array)-1 {
		temp := make([]int, len(array))
		copy(temp, array)
		*result = append(*result, temp)
		return
	}
	
	for j := index; j<len(array); j++ {
		swap(array, index, j)
		util(array, index+1, result)
		swap(array, index, j)
	}

}

func swap(array []int, i int, j int){
	array[i], array[j] = array[j], array[i]
}

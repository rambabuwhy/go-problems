package main

func ArrayOfProducts(array []int) []int {
	// Write your code here.
	
	product := make([]int, len(array))
	
	for i:= range array {
		product[i] = 1
	}
	
	leftProduct := 1
	for i := range array {
		product[i] = leftProduct
		leftProduct = leftProduct * array[i]
	}
	
	rightProduct := 1
	for i := len(array)-1; i>=0; i-- {
		product[i] = product[i] * rightProduct
		rightProduct = rightProduct * array[i]
	}
	
	return product
}

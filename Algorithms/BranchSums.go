package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	result := []int{}
	util (root, 0, &result)
	return result
}

func util(root *BinaryTree, currSum int, result *[]int){
	
	if root == nil {
		return
	}
	
	currSum = currSum + root.Value
	if root.Left == nil && root.Right == nil {
		*result = append(*result, currSum)
	}
	
	util(root.Left, currSum, result)
	util(root.Right, currSum, result)
	
}




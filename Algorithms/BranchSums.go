

/*

  Write a function that takes in a Binary Tree and returns a list of its branch
  sums ordered from leftmost branch sum to rightmost branch sum.
  A branch sum is the sum of all values in a Binary Tree branch. A Binary Tree
  branch is a path of nodes in a tree that starts at the root node and ends at
  any leaf node.
  
  
  */

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




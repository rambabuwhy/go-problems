package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	return util(root, 0)
}

func util(root *BinaryTree, depth int) int {
	if root == nil {
		return 0
	}
	return depth + util(root.Left, depth+1) + util(root.Right, depth+1)
}

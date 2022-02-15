/*
 Write a function that takes in a Binary Tree and inverts it. In other words,
  the function should swap every left node in the tree for its corresponding
  right node.
  
  */
  
package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	// Write your code here.
	Q := []*BinaryTree{}
	Q = append(Q, tree)
	
	for len(Q) > 0 {
		curr := Q[0]
		Q = Q[1:]
		
		curr.Left, curr.Right = curr.Right, curr.Left
		
		if curr.Left != nil {
			Q = append(Q, curr.Left)
		}
		
		if curr.Right != nil {
			Q = append(Q, curr.Right)
		}
		
	}
}

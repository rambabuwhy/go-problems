package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	
	curr := tree
	result := curr.Value
	for curr != nil {
		if abs(target, curr.Value) < abs(target, result){
			result = curr.Value
		}
		
		if target < curr.Value {
			curr = curr.Left
		} else if target > curr.Value {
			curr = curr.Right
		} else {
			break
		}
		
	}
	return result
}

func abs( i int, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

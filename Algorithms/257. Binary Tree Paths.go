/*
Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

 

Example 1:


Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
Example 2:

Input: root = [1]
Output: ["1"]
*/

func binaryTreePaths(root *TreeNode) []string {
    
    result := make([]string,0)
    util(root,"", &result)
    return result
    
}

func util(root *TreeNode, path string, result *[]string) {
    
    if root == nil{
        return
    }
    
    path = path + strconv.Itoa(root.Val)
    if root.Left == nil && root.Right == nil {
        *result = append(*result, path)
        return
    } 

    if root.Left != nil{
        util(root.Left, path + "->" , result)
    }
    
    if root.Right != nil {
        util(root.Right, path + "->" , result)
    }
}

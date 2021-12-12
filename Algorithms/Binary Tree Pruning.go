/*
Given the root of a binary tree, return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

A subtree of a node node is node plus every node that is a descendant of node.

*/

func pruneTree(root *TreeNode) *TreeNode {
    
    if root == nil {
        return nil
    }
    
    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    
    if root.Val == 0 && root.Left == nil && root.Right == nil {
        return nil
    }
    
    return root
}

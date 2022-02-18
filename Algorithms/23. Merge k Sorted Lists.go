/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

 

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    
    if len(lists) < 1{
        return nil
    }
    
    first := lists[0]
    for i:=1; i < len(lists); i++{
        first = merge(first, lists[i])
    }
    return first
}

func merge(l1,l2 *ListNode) *ListNode {
    
    dummy := &ListNode{}
    prev := dummy
    
    
    for l1 != nil && l2 != nil{
        
        if l1.Val <= l2.Val{
            prev.Next = l1
            l1 = l1.Next
        } else{
            prev.Next = l2
            l2 = l2.Next
        }
        
        prev = prev.Next
    }
    
    if l1 != nil {
        prev.Next = l1
    }
    
    if l2 != nil {
        prev.Next = l2
    }
    return dummy.Next  
}

/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

 

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.
*/


func intersection(nums1 []int, nums2 []int) []int {
    
    result := make(map[int]bool)
    
    m := make(map[int]bool)
    
    for i:=0; i<len(nums1); i++ {
        m[nums1[i]] = true
    }
    
    for i:=0; i<len(nums2); i++{
        if _,ok := m[nums2[i]]; ok {
            
            if _, k := result[nums2[i]]; !k{
                result[nums2[i]] = true
            }
           
        }
    }
    
    results := make([]int, 0)
    for k, _ := range result{
        results = append(results, k)
    }
    
    return results
    
}

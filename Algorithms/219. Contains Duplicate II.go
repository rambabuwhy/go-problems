/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

 

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

*/


func containsNearbyDuplicate(nums []int, k int) bool {
    hash := make(map[int]bool)
    for i:=0; i<len(nums); i++{
        if _,ok := hash[nums[i]];ok{
            return true
        }
            
        hash[nums[i]] = true
        if len(hash) > k {
            delete(hash, nums[i-k]) 
        }
    }
    return false
}

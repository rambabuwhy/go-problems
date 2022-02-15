/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

 

Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1
*/

//Remember XOR operation:
//a XOR 0 = a
//a XOR a = 0


func singleNumber(nums []int) int {
    
    if len(nums) < 1 {
        return 0
    }
    
    result := nums[0]
    for i:=1; i<len(nums); i++{
        result = result ^ nums[i]
    }
    return result
}

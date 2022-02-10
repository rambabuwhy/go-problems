/*
Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.

 

Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2

*/
func subarraySum(nums []int, k int) int {
    
    m := make(map[int]int)
    m[0]=1
    sum := 0
    count :=0
    
    for i := 0; i< len(nums); i++{
        sum = sum + nums[i]
        
        if _,ok := m[sum-k]; ok{
            count = count + m[sum-k];
        }
        if _, ok := m[sum]; ok  {
            m[sum]++;
        }else {
            m[sum]= 1;
        }
    }
    return count
}

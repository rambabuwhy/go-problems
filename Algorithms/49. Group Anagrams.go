/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]

*/

func groupAnagrams(strs []string) [][]string {
    
    result := [][]string{}
    m := make(map[string][]string)
    
    for _,s := range strs{
        
        origStr := s
        ss := []rune(s)
        sort.Slice(ss, func(a,b int) bool{
            return ss[a] < ss[b]
        })
        key := string(ss)
        m[key] = append(m[key], origStr)
               
    }
    
    for _, v := range m{
        result = append(result,v)
    }
    return result
    
}

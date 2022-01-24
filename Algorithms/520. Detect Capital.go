/*

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital, like "Google".
Given a string word, return true if the usage of capitals in it is right.

 

Example 1:

Input: word = "USA"
Output: true
Example 2:

Input: word = "FlaG"
Output: false

*/

import "unicode"
func detectCapitalUse(word string) bool {
   
    r := []rune(word)
    count := 0
    for _, r := range r  {
        if unicode.IsUpper(r) == true {
            count++
        }
        
    }
    
    if count ==1 && unicode.IsUpper(r[0]){
        return true
    }
    
    if count == 0 || count == len(r){
        return true
    }
    
    return false 
}

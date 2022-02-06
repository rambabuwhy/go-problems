
/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

 

Example 1:

Input: num1 = "11", num2 = "123"
Output: "134"
Example 2:

Input: num1 = "456", num2 = "77"
Output: "533"
Example 3:

Input: num1 = "0", num2 = "0"
Output: "0"
*/

func addStrings(num1 string, num2 string) string {
    
    
    l1 := len(num1)-1
    l2 := len(num2)-1
    
    x1 := 0
    x2 := 0
    carry := 0
    result := ""
    for l1>=0 || l2>=0  {
      
        if l1>=0 {
            x1,_ = strconv.Atoi(string(num1[l1]))
        } else{
            x1 = 0
        }
        
        if l2>=0 {
            x2,_ = strconv.Atoi(string(num2[l2]))
        }else {
            x2 = 0
        }
        
        val := (x1 + x2 + carry) % 10
        carry =  (x1 + x2 +carry) / 10
        result = strconv.Itoa(val) + result
        
        l1--
        l2--
    }
    
    if carry != 0 {
        result = strconv.Itoa(carry) + result
    }
    return result
}

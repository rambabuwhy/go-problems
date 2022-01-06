/*

The Tribonacci sequence Tn is defined as follows: 

T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.

 

Example 1:

Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
Example 2:

Input: n = 25
Output: 1389537

*/

func tribonacci(n int) int {
    
    
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }else if n == 2 {
        return 1
    }
    
    first := 0
    second := 1;
    third := 1;
    fourth := first + second + third 
    
    for i := 3; i<=n; i++{
        fourth = first + second + third
        first = second
        second = third
        third = fourth
    }
    return fourth

}

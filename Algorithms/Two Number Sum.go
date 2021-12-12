func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	
	result := []int{}
	lmap := map[int]bool{}
	for _,val := range array {
		second := target - val
		if _, found := lmap[second]; found{
			result = append(result,val)
			result = append(result, second)
		}
		lmap[val] = true
		
	}
	return result
}

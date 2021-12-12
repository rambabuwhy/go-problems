package main

func ReverseWordsInString(str string) string {
	// Write your code here.
	
	//convert string to chars array
	var cstr []byte
	cstr = make([]byte, 0)
	
	for _,c := range []byte(str) {
		cstr = append(cstr, c)
	}
	
	//reverse whole string
	reverse(cstr, 0, len(str)-1)
	
	//revsese word by word
	start := 0
	for start < len(cstr) {
		end := start
		for end < len(cstr) && cstr[end] != ' ' {
			end = end + 1
		}
		
		reverse(cstr, start, end-1)
		start = end + 1
	}
	
	return string(cstr)
}

//reverse with given start and end index
func reverse(str []byte, start int, end int){
	
	for start < end {
		str[start], str[end] = str[end], str[start]
		start = start + 1
		end = end -1
	}
	
}

package main

import "fmt"

func BubbleSort(s []int) {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10-i-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j)
			}
		}
	}

}

func swap(s []int, i int) {

	temp := s[i+1]
	s[i+1] = s[i]
	s[i] = temp

}

func main() {

	var i int
	s := make([]int, 0, 10)

	for j := 0; j < 10; j++ {

		fmt.Scan(&i)

		s = append(s, i)

	}

	fmt.Println("Before:", s)

	BubbleSort(s)

	fmt.Println("After:", s)

}

package main

import "fmt"

func main() {
	i := 1
	defer fmt.Println("defer:", i)
	i++
	i = 10
	fmt.Println("end")
}

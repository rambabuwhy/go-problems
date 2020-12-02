package main

import "fmt"

func main() {
	x := 100
	fmt.Println(x)

	y := x + 100
	fmt.Println(y)

	fmt.Printf("%T  \n", y)

	//string
	var s string
	s = "ram"
	fmt.Println("string s:", s, "adding")

	//integer
	var i int
	var j int = 100
	fmt.Println("integer i:", i, "adding")
	fmt.Println("integer j:", j, "adding")

	foo()

}

func foo() {
	fmt.Println("foo  func")
}

package main

import (
	"fmt"
)

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

	//fmt

	var a int = 10
	fmt.Printf("%v  \n", a)

	//own type

	type T int
	var b T = 400
	fmt.Printf("%T \n", b)

	//convert it to int

	a = int(b)
	fmt.Printf("%v  \n", a)

	//bool

	var c bool
	c = true
	fmt.Println(c)

	foo()

	//types

	var p uint8
	p = 23
	fmt.Println(p)

	//bytes

	var str string = " Rambabu Yerajana"

	var B []byte = []byte(str)

	fmt.Println(B)

	for i2 := 0; i2 < len(str); i2++ {
		fmt.Printf("%#U \n", str[i2])
	}

}

func foo() {
	fmt.Println("foo  func")
}

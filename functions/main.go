package main

import "fmt"

func main() {

	defer foo()

	boo()

	car()

}

func foo() {

	fmt.Println("foo")

}

func boo() {

	fmt.Println("boo")

}

func car() {

	fmt.Println("car")

}

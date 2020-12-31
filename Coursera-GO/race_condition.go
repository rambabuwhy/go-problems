package main

import (
	"fmt"
)

/*

Race condition:  two methods add() and sub()
for adding/subtracting the same variable
Due to the uncertainty of Goroutine scheduling mechanism,
the results of the following program is unpredictable.

*/
func add(i *int) {
	fmt.Println("add:", *i)
	*i = *i + 1
}

func sub(i *int) {
	*i = *i - 1
	fmt.Println("sub i:", *i)

}

func main() {

	var i int = 0
	for {

		go add(&i)
		go sub(&i)

	}

}

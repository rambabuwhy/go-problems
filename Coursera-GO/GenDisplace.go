package main

import "fmt"

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(t float64) float64 {

	//s := (a*(t*t))/2 + v0*t + s0
	f1 := func(t float64) float64 {
		//var s float64
		var a1 float64
		var a2 float64
		var a3 float64
		a1 = (a * (t * t)) / 2
		a2 = v0 * t
		a3 = s0

		return float64(a1 + a2 + a3)
	}
	return f1
}

func main() {

	var a, v0, s0, t float64
	fmt.Println("Enter a v0 s0:")
	fmt.Scan(&a)
	fmt.Scan(&v0)
	fmt.Scan(&s0)
	fmt.Println("Enter time:")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println("ans:", fn(t))
}

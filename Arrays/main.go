package main

import "fmt"

func main() {

	var y []int

	y = []int{4, 5, 6, 7, 8}

	fmt.Println(y)

	for i, v := range y {
		fmt.Println(i, v)
	}

	x := []int{1, 2}

	z := append(x, y...)

	fmt.Println(x)
	fmt.Println(z)

	m := map[string]int{

		"ram":  20,
		"babu": 30,
	}

	fmt.Println(m)

	fmt.Println(m["ram"])

	//m.append("avy",100)

	m["babue"] = 100

	v, err := m["babue"]

	fmt.Println(v, err)
}

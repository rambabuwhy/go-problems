package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var i string

	slice := make([]int, 3)

	L := len(slice)

	for {

		fmt.Scan(&i)

		d, _ := strconv.Atoi(i)

		if i == "X" {
			break
		} else {

			if L > 0 {
				slice[L-1] = d
				L--
			} else {
				slice = append(slice, d)
			}
			sort.Ints(slice)

			fmt.Println(slice)

		}

	}

}

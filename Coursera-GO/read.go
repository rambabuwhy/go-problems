package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string `valid:MaxSize(20)`
	lname string `valid:MaxSize(20)`
}

func main() {

	var result []person
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("err:", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		var p person

		value := strings.Fields(scanner.Text())

		if len(value[0]) <= 20 {
			p.fname = value[0]
		} else {
			p.fname = value[0][0:20]
		}

		if len(value[1]) <= 20 {
			p.lname = value[1]
		} else {
			p.lname = value[1][0:20]
		}

		result = append(result, p)
	}

	fmt.Println(result)

	file.Close()

}

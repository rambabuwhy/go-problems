package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := make(map[string]string)
	var name string
	var address string

	fmt.Scan(&name)
	fmt.Scan(&address)

	m["name"] = name
	m["address"] = address

	j, err := json.Marshal(m)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(j)         //byte  format
	fmt.Println(string(j)) //string format

}

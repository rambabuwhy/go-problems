package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println(len(text))

	/*
		if (s[0] == 'i' || s[0] == 'I') &&
			(s[len(s)-1] == 'n' || s[len(s)-1] == 'N') &&
			(strings.Contains(s, "a") || strings.Contains(s, "A")) {
			fmt.Println("Found!")
		} else {

			fmt.Println("Not Found!")
		}

	*/

}

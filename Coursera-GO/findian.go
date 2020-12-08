package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	if (s[0] == 'i' || s[0] == 'I') && (s[len(s)-1] == 'n' || s[len(s)-1] == 'N') && (strings.Contains(s, "a") || strings.Contains(s, "A")) {
		fmt.Println("Found!")
	} else {

		fmt.Println("Not Found!")
	}


}

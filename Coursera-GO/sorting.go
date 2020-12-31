package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Print("input with space > ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	s := strings.Split(strings.TrimSpace(input), " ")
	a := make([]int, 0, len(s))

	for _, str := range s {
		num, _ := strconv.Atoi(str)
		a = append(a, num)
	}

	const par = 4
	n := int(math.Max(math.Ceil(float64(len(s))/float64(par)), 1))

	var w sync.WaitGroup

	for i := 0; i < par; i++ {
		from := int(math.Min(float64(i*n), float64(len(a))))
		to := int(math.Min(float64((i+1)*n), float64(len(a))))

		w.Add(1)

		go func(arr []int) {
			fmt.Printf("sort: %v \n", arr)

			sort.Ints(arr)

			w.Done()
		}(a[from:to])
	}

	w.Wait()

	sort.Ints(a)
	fmt.Println("\nresult: ", a)
}

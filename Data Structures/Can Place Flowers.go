package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	i := 0
	count := 0

	for i < len(flowerbed) {

		//three conditions  1. ith position is

		c1 := flowerbed[i] == 0                              //  1.  It should be 0
		c2 := (i == 0 || flowerbed[i-1] == 0)                // 2.  previous position should be 0
		c3 := (i == len(flowerbed)-1 || flowerbed[i+1] == 0) // 3.   next position should be 0

		if c1 && c2 && c3 {

			flowerbed[i] = 1
			count = count + 1
		}

		i = i + 1
	}

	return count >= n

}

func main() {
	var flowerbed = make([]int, 0, 0)
	canPlaceFlowers(flowerbed, 2)
}

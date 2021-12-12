package main
import "sort"
func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	// Write your code here.
	Red := redShirtHeights
	Blue := blueShirtHeights
	sort.Slice(Red, func(i int, j int) bool {
		return Red[i] > Red[j]
	})
	sort.Slice(Blue, func(i int, j int) bool {
		return Blue[i] > Blue[j]
	})
	
	firstRowColour := "Blue"
	if Red[0] < Blue[0] {
		firstRowColour = "Red"
	}
	
	for i, _ := range redShirtHeights {
		if firstRowColour == "Red"{
			if Red[i] >= Blue[i] {
				return false
			}
		}	else {
			if Blue[i] >= Red[i] {
				return false
			}
		}
	}
	return true
}

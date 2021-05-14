package task

import (
	"math/rand"
	"sort"
	"time"
)

func BinarySearch(array []string, toFind string) int {
	//fmt.Print("\n\n")
	rand.Seed(time.Now().UnixNano())
	sort.Strings(array)
	/*fmt.Print("Array: [")
	for i, elem := range array {
		if i != len(array)-1 {
			fmt.Print("\"", elem, "\", ")
		} else {
			fmt.Print("\"", elem, "\"")
		}
	}
	fmt.Print("]\n")
	fmt.Print("Text: \"" + toFind + "\"")*/
	left := 0
	right := len(array)
	mid := (right + left) / 2
	for {
		if array[mid] == toFind {
			break
		}
		if toFind > array[mid] {
			left = mid
		} else {
			right = mid
		}
		mid = (right + left) / 2
	}
	return mid
}

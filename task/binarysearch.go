package task

import (
	"fmt"
	"math/rand"
	"time"
)

func Sort(array []string) []string {
	for u := 0; u < len(array)-1; u++ {
		for i := u + 1; i < len(array); i++ {
			if array[i] < array[u] {
				array[i], array[u] = array[u], array[i]
			}
		}
	}
	return array
}

func BinarySearch() {
	fmt.Print("\n\n")
	rand.Seed(time.Now().UnixNano())
	array := Sort([]string{"Ubuntu", "Linux Mint", "OpenSUSE", "Debian", "Arch Linux", "Deepin OS", "Elementary OS", "Manjaro Linux", "Fedora", "Kali Linux"})
	fmt.Print("Array: [")
	for i, elem := range array {
		if i != len(array)-1 {
			fmt.Print("\"", elem, "\", ")
		} else {
			fmt.Print("\"", elem, "\"")
		}
	}
	fmt.Print("]\n")
	toFind := array[rand.Intn(len(array))]
	fmt.Println("Text: \"" + toFind + "\"")
	left := 0
	right := len(array)
	mid := (right + left) / 2
	for {
		if array[mid] == toFind {
			fmt.Println("Id: ", mid)
			break
		}
		if toFind > array[mid] {
			left = mid
		} else {
			right = mid
		}
		mid = (right + left) / 2
	}
	fmt.Print("\n\n")
}

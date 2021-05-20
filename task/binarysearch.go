package task

import (
	"math/rand"
	"sort"
	"time"
)

func BinarySearch(array []string, toFind string) int {
	rand.Seed(time.Now().UnixNano())
	sort.Strings(array)
	left := 0
	right := len(array)
	mid := (right + left) / 2
	for ; ; mid = (right + left) / 2 {
		if array[mid] == toFind {
			break
		} else if toFind > array[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	return mid
}

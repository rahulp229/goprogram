package main

import (
	"fmt"
	"sort"
)

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2
		fmt.Println("low = ", low)
		fmt.Println("high = ", high)
		fmt.Println("median = ", median)

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}

	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 3}
	sort.Slice(items, func(i, j int) bool { return items[i] < items[j] })
	fmt.Println(binarySearch(3, items))

}

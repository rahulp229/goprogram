package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{90, -20, 40, 75, -70, 60, 30, -10, -80, -50}
	a2 := []int{}
	a3 := []int{}
	result := []int{}
	//num1 := sort.RE(num)
	//
	//var result1 string
	fmt.Println("len of a : ", len(a))

	a2 = a
	a3 = a
	sort.Slice(a, func(i, j int) bool { return a2[i] < a2[j] })

	//sort.Slice(a, func(i, j int) bool { return a3[i] > a3[j] })
	sort.Sort(sort.Reverse(sort.IntSlice(a3)))
	fmt.Println("asending : ", a2)

	fmt.Println("desending : ", a3)

	for i := 0; i < len(a3); i++ {
		result = append(result, a3[i], a2[i])
	}
	fmt.Println("result : ", result)
	fmt.Println("original : ", a)
	a1 := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Remove the element at index i from a.
	a1[i] = a1[len(a1)-1] // Copy last element to index i.
	a1[len(a1)-1] = ""    // Erase last element (write zero value).
	a1 = a1[:len(a1)-1]   // Truncate slice.

	//fmt.Println(a1) // [A B E D]

}

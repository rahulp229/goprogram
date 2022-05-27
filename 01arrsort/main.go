package main

import (
	"fmt"
	"sort"
)

// write a program to sort slice in asending and desending order with and without sort.Slice function
func main() {
	var arr = []int{1, 34, 56, 3, 78}
	var arr2 = []int{1, 34, 56, 3, 78}
	n := len(arr)
	var arr1 = []int{}
	//var arr1 = []int{2}
	temp := 0
	for i := 0; i < n; i++ {

		arr1 = append(arr1, arr[i])

		temp = temp + 1
		fmt.Println(arr[i])

	}
	fmt.Println(arr1)

	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })
	fmt.Println("asending - ", arr2)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println("sort.Sort - sort.Reverse - ", arr)

	arr3 := []float64{34.4, 13.3, 3.2, 45.6}
	sort.Float64s(arr3)
	fmt.Println(arr3)
	x := 1
	fmt.Println("arrr2", arr2)
	i := sort.Search(len(arr2), func(i int) bool { return arr2[i] == x })
	if i < len(arr2) && arr2[i] == x {
		// x is present at data[i]
		fmt.Println("present", i)
	} else {
		// x is not present in data,
		// but i is the index where it would be inserted.
		fmt.Println("not present", i)

	}
	fmt.Println(i)
	names := []string{"Aadi", "Deepali", "Rajani", "Kishor", "Rahul", "Rahul"}
	i = sort.SearchStrings(names, "Rahul")
	//fmt.Println(names[i])
	if len(names) > i {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}

	sort.Strings(names)
	fmt.Println("strings >>> ", names)
	// Remove duplicate
	m := map[string]bool{}
	mm := []string{}
	for _, val := range names {
		if m[val] == false {
			m[val] = true
			mm = append(mm, val)
		}
	}
	fmt.Println(mm)
	names = mm
	fmt.Println(names)

}

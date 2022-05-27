package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\nQ. How many ways to sort array/slice in Golang?")
	str := ` 
Answer :- There are three ways we can sort array/slice
// input array
num := []int{123, 32, 563, 46, 5124, 26}

1. Using for loop reverse the array
for i, j := len(num)-1, 0; i > j; i, j = i-1, j+1 {
	temp := num[j]
	num[j] = num[i]
	num[i] = temp
}
fmt.Println("reverse : ", num)
// output : [26 5124 46 563 32 123]
-------------------------------------------------------------------
2. Using sort.Slice in asending order

sort.Slice(num, func(i, j int) bool { return num[i] < num[j] }) 
// note :- '<'(lt) : for asending , '>'(gt) : for desending 
fmt.Println("asending : ", num)
// output : [26 32 46 123 563 5124]
---------------------------------------------------------------------
3. Using sort.Sort in desending order

sort.Sort(sort.Reverse(sort.IntSlice(num)))
fmt.Println("desending : ", num)
// output : [5124 563 123 46 32 26]
`
	//	fmt.Println(len(str))
	//	fmt.Println(string(str[48]))
	for i := 0; i < len(str); i++ {
		time.Sleep(140 * time.Millisecond)
		fmt.Printf("%v", string(str[i]))
	}
}

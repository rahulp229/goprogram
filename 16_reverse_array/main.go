package main

import (
	"fmt"
)

func main() {

	num1, num2 := 10, 13
	temp := num1
	num1 = num2
	num2 = temp

	fmt.Println("Number 1 = ", num1, " number2 = ", num2)

	//a := []int{6676, 3216, 4063, 8373, 423, 586, 8850, 6762}
	//   6676  3216  4063  8373  423  586  8850  6762
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var prev int
	for key, val := range a {
		fmt.Println("key ", key, " prev ", prev)

		//prev := (len(a) - 1) - key
		if (key >= prev) && (key != 0) {
			break
		}
		prev = (len(a) - 1) - key //7
		current := val
		a[key] = a[(len(a)-1)-key]
		a[(len(a)-1)-key] = current
	}
	fmt.Println("a ", a)
	//key0 ->	prev = 0
	//key1 -> prev = 4
	//key2 -> prev = 3

}

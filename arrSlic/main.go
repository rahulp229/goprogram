package main

import "fmt"

func main() {
	//initialise array
	var number [5]int

	number[0] = 10
	number[1] = 20
	number[2] = 30
	number[3] = 40
	number[4] = 50

	fmt.Println("array : ", number)

	number2 := number[:3] // fetch elements from index 0 to 3

	fmt.Println("range 0- 3 :", number2)

	number3 := number[3:] // fetch elements from index 3 to till last element

	fmt.Println("range 3 to end element : ", number3)

	number4 := number[1:4] // fetch elements  from index 1 to 4

	fmt.Println("range 1 to 4 : ", number4)

	fmt.Println("thank you for watching this video if you like then hit like button and subscribe my channel......")
}

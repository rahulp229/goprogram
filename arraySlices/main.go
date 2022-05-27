package main

import (
	"fmt"
	"reflect"
)

type Details struct {
	Test11 string
}

type Details2 struct {
	Test11 string
}

type Elem struct {
	Details interface{}
}

func main() {
	fmt.Println("Hello, playground")
	numbers := []int{12, 12, 23, 45, 12}
	//fmt.Println("numbers >>>>>",numbers)
	dd1 := Details{Test11: "qwerty"}
	elem11 := Elem{Details: dd1}
	fmt.Println(">>>>>>", reflect.TypeOf(elem11.Details.Test11))
	numMap := map[int]bool{}
	testmap := map[int]int{}
	var count int
	for _, val := range numbers {
		if numMap[val] == true {
			//fmt.Print("11")
			testmap[val] = testmap[val] + 1
		} else {
			numMap[val] = true
			testmap[val] = count + 1
		}
	}
	fmt.Println("final", testmap)
	c := make(chan string, 2)
	c <- "10"
	fmt.Println(<-c)
}

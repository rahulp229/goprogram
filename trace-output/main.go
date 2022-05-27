package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	colors := [4]string{"red", "green", "blue", "pink"}

	numColors := colors[2:3]

	fmt.Println(numColors, len(numColors), cap(numColors)) // guess the output

	numColors = append(numColors, "black")
	fmt.Println(colors, len(colors), cap(colors)) // guess the output

	names := [3]string{"a", "b", "c"}

	sl := names[1:]
	fmt.Println(sl, len(sl), cap(sl)) // guess the output ["a"] 1, 2
	//==========================================
	colorsNew := [5]string{"red", "green", "blue", "pink", "purple"}

	newSlic := colorsNew[1:2]

	fmt.Println("new slice, cap", newSlic, cap(newSlic), len(newSlic))
	newSlic = append(newSlic, "A", "B", "C")
	fmt.Println(newSlic)
	fmt.Println(colorsNew)

	//country := map[string]string{"india": "delhi"}

	//ok := A("test1")
	//mt.Println("erros", ok.Error())
	// if ok != nil {
	// 	err := fmt.Errorf("soem error %s ", ok.Error())
	// 	fmt.Printf("error : %v", err)
	// 	///fmt.Errorf("soem error %v ", ok.Error())
	// }

	//fmt.Println("success")

	//some code ===========================================
	num := []int{123, 32, 563, 46, 5124, 26}
	fmt.Println("before : ", num)
	for i, j := len(num)-1, 0; i > j; i, j = i-1, j+1 {
		temp := num[j]
		num[j] = num[i]
		num[i] = temp
	}
	fmt.Println("reverse : ", num)
	sort.Slice(num, func(i, j int) bool { return num[i] < num[j] })
	fmt.Println("asending : ", num)
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println("desending : ", num)

//====================================================

}

func A(input string) error {

	if input != "test" {
		return errors.New("some error")
	}
	return nil
}

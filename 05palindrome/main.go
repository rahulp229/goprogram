package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter string: ")
	var reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("input : ", input)
	flag := IsPalindrom(input)
	if flag {
		fmt.Println("string is palidrome", input)
		os.Exit(1)
	}
	fmt.Println("string is not palidrome", input)
}

func IsPalindrom(input string) bool {
	//in1 := "madam"
	fmt.Println("before = ", input)
	input = strings.ToLower(input)
	fmt.Println("after = ", input)
	for i, j := len(strings.Split(input, "\n")[0])-1, 0; i >= j; i, j = i-1, j+1 {
		first := string(input[i])
		second := string(input[j])
		// fmt.Println(" first value", first)
		// fmt.Println(" i value", i)
		// fmt.Println(" second value", second)
		// fmt.Println(" j value", j)
		if first != second {
			// fmt.Println(" i value", first, i)
			// fmt.Println(" j value", second, j)
			return false
		}
	}
	return true
}

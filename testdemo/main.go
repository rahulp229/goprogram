package main

import "fmt"

func calculate(x int) int {
	result := x + 2
	return result
}

var Braces map[string]string

func main() {
	out := calculate(2)
	fmt.Print("out", out)
	Braces["("] =")"
	Braces["["] ="]"

}

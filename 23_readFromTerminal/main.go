package main

import "fmt"

func main() {
	n := 121
	for i := 0; i < 3; i++ {
		rem := n % 10
		n = n / 10
		fmt.Println("rem ", rem)
		fmt.Println("n", n)
	}

}

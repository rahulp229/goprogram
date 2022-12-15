package main

import "fmt"

func main() {
	numbers := []int{2, 5, 3, 7, 5, 9, 13}
	fmt.Println("sum - ", sumOfSquare(numbers))
}

func sumOfSquare(numbers []int) int {
	sum := 0
	go func() {

	}()
	for i := 0; i < len(numbers); i++ {

		sqr := numbers[i] * numbers[i]
		sum += sqr
	}
	return sum
}

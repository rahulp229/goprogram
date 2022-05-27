package main

import "fmt"

func main() {
	result := []int{}
	numbers := []int{20, 10, 40, 30, 60, 50}

	for k := 0; k < len(numbers); k++ {
		for i, j := 0, 1; j < len(numbers); i, j = i+1, j+1 {
			if numbers[i] > numbers[j] {
				temp := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = temp
			}
		}

	}
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("large = %v and Small = %v \n", numbers[j], numbers[i])
		result = append(result, numbers[j], numbers[i])
	}
	fmt.Println(numbers)
	fmt.Println(result)
}

// 1,2,3,4,5,6
// 6 1 i,j=0,5
// 6   i,j=1,4
// 	i,j=2,3
// 	i,j=3,2
// 	i,j=4,1
// 	i,j=5,0

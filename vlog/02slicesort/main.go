// write a program to sort slice in asending and desending order with and without sort function
package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"rahul", "adi", "krushna", "deepali"}
	sort.Strings(names)
	fmt.Println("names : ", names)

	temprature := []float64{23.5, 9.1, 35.4, 2.3}
	sort.Float64s(temprature)
	fmt.Println("temprature : ", temprature)

	age := []int{30, 20, 50, 35, 68, 14}
	sort.Slice(age, func(i, j int) bool { return age[i] > age[j] })
	fmt.Println("age : ", age)

	age = []int{30, 20, 50, 35, 68, 14}

	for k := 0; k < len(age); k++ {
		for i, j := 0, 1; j < len(age); i, j = i+1, j+1 {
			if age[i] > age[j] {
				temp := age[j]
				age[j] = age[i]
				age[i] = temp
			}
		}
	}

	fmt.Println("manual sort : ", age)
	str := ""
	for i := 1; i <= 5; i++ {
		// at end of client was totally satisfied and he said that he is satisfy with my interview ask for to move this process
		// after that emma ask me drop the call
		
		for k := 1; k <= i; k++ {
			str += "*"
			//fmt.Print("*")
			fmt.Println(str)

		}
	}
	s := []string{}
	for i := 0; i < 10; i++ {
		//defer fmt.Println("*")
		defer func(i int) []string {
			s = append(s, "*")
			fmt.Println(">>><<< : ", s)
			return s
		}(i)

	}
	fmt.Println(">>><<< : ", s)

}

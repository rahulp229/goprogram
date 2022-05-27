package main

import "fmt"

func main() {
	names := []string{"rahul", "virat", "rohit"}

	//iterating over names

	for key, val := range names {
		fmt.Println("index : ", key)
		fmt.Println("value: ", val)
	}

	// range for iterating maps

	cc := map[string]string{"india": "delhi", "USA": "washington"}
	fmt.Println("------------------------")
	for index, value := range cc {
		fmt.Println("country : ", index)
		fmt.Println("capital: ", value)
	}

	for index1 := range cc {
		fmt.Println("country : ", index1)
	}

	test := [6]int{50, 10, 60, 70, 30, 80}
	fmt.Println("cap :", cap(test))

	for  j := 1; j < len(test); j++ {
		if test[0] < test[j]{
			test[0] = test[j]
		}
	}

}

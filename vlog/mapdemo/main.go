package main

import "fmt"

func main() {
	//there are two ways to declar maps
	//1st
	var country = map[string]string{}

	country["India"] = "New Delhi"
	country["US"] = "New York"

	fmt.Println("country ==> ", country)

	//2nd
	var user = make(map[string]int)
	user["rahul"] = 23423423
	user["rakesh"] = 564545
	fmt.Println("User data ", user)

	// short hand for maps

	colors := map[string]string{"Red": "#345345", "Yellow": "#FFF"}
	fmt.Println("colors", colors)

	// iterating over maps
	for key, value := range country {
		fmt.Printf("%s is capital of %s", value, key)
		fmt.Println("")
	}

}

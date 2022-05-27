package main

import "fmt"

func main() {
	var testVar interface{} = 3.40282346638528859811704183484516925440e+38

	switch testVar.(type) {
	case int:
		fmt.Println("this int type value")
	case string:
		fmt.Println("this is string")
	case float64:
		fmt.Println("this type float64")

	}
}

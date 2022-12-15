package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// perfect square input 3
	number := 3
	flag := isPerfectSquare(float64(number))
	if flag {
		fmt.Println("it is perfect square", flag)

	} else {
		fmt.Println("it is not perfect square", flag)
	}

}
func isPerfectSquare(n float64) bool {
	sr := math.Sqrt(n)
	s := strconv.FormatFloat(sr, 'g', 2, 64)
	s2 := strings.Split(s, ".")
	fmt.Println("value : ", s)
	//fmt.Println("debug >>> : ", s2[1])
	if len(s2) == 1 {
		return true
	}

	return false
}

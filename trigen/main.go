package main

import (
	"fmt"
	"sort"
)

var NoofTestCases = 2

type TestCase struct {
	NoofFriends     int
	NoofKg          int
	PricePerPackets []int
}
type TestCases []TestCase

func main() {

	//	noofFriends := 3
	//	noofKg := 5

	//pricePerPackets := []int{-1, -1, 3, 5, -1}
	tc := TestCases{
		{NoofFriends: 3, NoofKg: 5, PricePerPackets: []int{-1, -1, 3, 5, -1}},
		{NoofFriends: 5, NoofKg: 4, PricePerPackets: []int{1, 2, 3, 4, 5}},
	}
	//noofTestCases := len(tc)
	flag := false
	for _, val := range tc {
		sort.Slice(val.PricePerPackets, func(i, j int) bool { return i < j })
		for i := 0; i < len(val.PricePerPackets); i++ {
			if val.PricePerPackets[i] > 0 {
				if i+1 == val.NoofKg {
					fmt.Println(val.PricePerPackets[i])
					flag = true
					break
				}

			}
			if (len(val.PricePerPackets)-2 == i) && (flag == false) {
				fmt.Println("-1")
			}
		}

	}

}

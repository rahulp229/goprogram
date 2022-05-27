package main

import (
	"fmt"
	"sort"
)

func main() {
	user := map[string]int{
		"Kishor": 41,
		"Aadi":   3,
		"Dipu":   28,
		"Rahul":  31,
		"Rajani": 48,
	}
	s := []int{}
	fmt.Println("befor : ", user)
	for _, val := range user {
		s = append(s, val)
	}

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println("map key are sorted : ", s, user)
}

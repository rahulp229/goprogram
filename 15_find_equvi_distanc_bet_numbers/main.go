package main

import (
	"fmt"
	"math"

	"google.golang.org/protobuf/internal/errors"
)

func main() {
	fmt.Println("Hello, playground")
	//arr := []float64{6, 7, 9, 11}
	result := []float64{}
	errors.Wrap()
	arr := []float64{6, 8, 10, 14, 18, 14, 43}
	for key, val := range arr {
		if key < len(arr)-1 {

			diff := math.Abs(val - arr[key+1])
			//mid = math.Abs(diff)
			mid := diff / 2

			fmt.Println("diff ", diff)

			if diff > 2 {
				if val < arr[key+1] {
					result = append(result, val+mid)
				} else {
					result = append(result, val-mid)
				}
			}
		}
	}

	fmt.Println("result ", result)

}

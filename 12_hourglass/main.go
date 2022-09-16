package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var twodArr = []string{"1 1 1 0 0 0", "0 1 0 0 0 0", "1 1 1 0 0 0", "0 0 2 4 4 0", "0 0 0 2 0 0", "0 0 1 2 4 0"}

func main() {
	numArr := [][]int{}

	for i := 0; i < len(twodArr); i++ {
		arrRowTemp := strings.Split(twodArr[i], " ")
		temArr := []int{}
		for _, item := range arrRowTemp {
			num, _ := strconv.ParseInt(item, 10, 64)
			temArr = append(temArr, int(num))
		}
		numArr = append(numArr, temArr)
	}
	hg := [][]int{}
	temp := []int{}
	for key, value := range numArr {
		if key == 4 {
			os.Exit(0)

		}
		item1 := numArr[key+1]
		item2 := numArr[key+2]

		for i := 0; i < 4; i++ {
			temp = append(temp, value[i], value[i+1], value[i+2])
		}
	}
	fmt.Println(numArr)
}

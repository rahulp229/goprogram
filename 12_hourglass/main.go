package main

import (
	"fmt"
	"strconv"
	"strings"
)

//var twodArr = []string{"1 1 1 0 0 0", "0 1 0 0 0 0", "1 1 1 0 0 0", "0 0 2 4 4 0", "0 0 0 2 0 0", "0 0 1 2 4 0"}
//var twodArr = []string{"-9 -9 -9 1 1 1", "0 -9 0 4 3 2", "-9 -9 -9 1 2 3", "0 0 8 6 6 0", "0 0 0 -2 0 0", "0 0 1 2 4 0"}
var twodArr = []string{"-1 -1 0 -9 -2 -2",
	"-2 -1 -6 -8 -2 -5",
	"-1 -1 -1 -2 -3 -4",
	"-1 -9 -2 -4 -4 -5",
	"-7 -3 -3 -2 -9 -9",
	"-1 -3 -1 -2 -4 -5",
}

func main() {
	numArr := [][]int64{}

	for i := 0; i < len(twodArr); i++ {
		arrRow := strings.Split(twodArr[i], " ")
		temArr := []int64{}
		for _, item := range arrRow {
			num, _ := strconv.ParseInt(item, 10, 64)
			temArr = append(temArr, num)
		}
		numArr = append(numArr, temArr)
	}
	hg := [][]int64{}
	result := map[int64][]int64{}
	fmt.Println("num arr ", numArr)
	var sum, tempSum int64
	for key, value := range numArr {
		if key == 4 {
			break
		}
		item1 := numArr[key+1]
		item2 := numArr[key+2]

		for i := 0; i < 4; i++ {
			temp := []int64{}
			tempSum = 0
			tempSum = value[i] + value[i+1] + value[i+2] + item1[i+1] + item2[i] + item2[i+1] + item2[i+2]
			temp = append(temp, value[i], value[i+1], value[i+2])
			temp = append(temp, item1[i+1])
			temp = append(temp, item2[i], item2[i+1], item2[i+2])
			//	fmt.Println(i, "temp ", temp)
			//fmt.Println("sum ", sum, " tempsum ", tempSum)
			if tempSum >= sum {
				//	fmt.Println("before sum ", sum)
				result = map[int64][]int64{}
				sum = tempSum
				//fmt.Println("tempsum ", tempSum, " sum ", sum)
				result[sum] = temp
			}
			hg = append(hg, temp)
		}
	}
	//fmt.Println("hourglass ", hg)
	fmt.Println("result ", result)
}

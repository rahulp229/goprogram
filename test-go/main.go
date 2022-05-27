package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//remove duplicate element from array

func RemoveElem(country []string) []string {
	temp := map[string]bool{}
	finalArray := []string{}
	for i, val := range country {
		fmt.Println(">>>>", temp[val])
		fmt.Printf("result %d= %v \n", i, finalArray)
		if temp[val] == true {
			fmt.Errorf("already present")
		} else {
			temp[val] = true
			finalArray = append(finalArray, val)
		}
	}

	return finalArray
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, er := reader.ReadString('\n')
	if er == io.EOF {
		fmt.Errorf("err %v", er)
	}
	fmt.Scan()
	fmt.Println("+++++", s)
	if strings.Split(s, "\n")[0] == "rahul" {
		fmt.Println("found")
	}
	var numbers []string
	for i := 0; i < 5; i++ {
		numbersItemTemp, _, err := reader.ReadLine()
		checkError(err)
		m, _ := strconv.ParseInt(string(numbersItemTemp), 0, 8)
		fmt.Println(">>>", m)
		numbersItem := fmt.Sprintf("%s", numbersItemTemp)
		numbers = append(numbers, numbersItem)
	}
	fmt.Println("-----", numbers)

	country := []string{"india1", "india", "england"}

	result := RemoveElem(country)

	fmt.Println("result = ", result)
	fmt.Println("result22 = ", country[:2-1])
	fmt.Println("result23= ", country[2:])
	ss := []int{}
	fmt.Println("result24= ", cap(ss))
	var tt interface{} = 11
	switch tt.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

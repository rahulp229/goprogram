package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'closestNumbers' function below.
 *
 * The function accepts INTEGER_ARRAY numbers as parameter.
 */

func closestNumbers(numbers []int32) {
    // Write your code here
    fmt.Println("----")
    tempArr := [][]int32{}
    var diff, tempDiff, k1, k2  int32 = 0, 0, 0, 0
    for k, _ := range numbers{
        for i, _ := range numbers{
            if k != i && k < i{
                tempDiff = numbers[k] - numbers[i] 
                if tempDiff < diff && tempDiff != 0 {
                    k1 = numbers[k]
                    k2 = numbers[i]
                    diff= tempDiff                     
                }
                // 4 - 4 = 0.  1
                // 4 - 2 = 2.  2 d
                //4 - 1 =  3      3
                //4 - 3 = 1   4  d     
            }
        }
        tt := []int32{k1, k2}
        tempArr = append(tempArr, tt)
    }
    fmt.Println("===", tempArr)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    numbersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var numbers []int32
fmt.Println(">>>>",numbersCount)
    for i := 0; i < int(numbersCount); i++ {
        numbersItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        numbersItem := int32(numbersItemTemp)
        numbers = append(numbers, numbersItem)
    }

    closestNumbers(numbers)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

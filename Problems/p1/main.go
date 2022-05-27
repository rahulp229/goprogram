/*
 *Read input from STDIN. Print your output to STDOUT
 *Use fmt.Scanf to read input from STDIN and fmt. Println to write output to STDOUT
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	fmt.Print("Virun composition : ")
	virusComposition := strings.TrimSpace(readLine(reader))
	fmt.Println("")

	fmt.Print("No of People : ")
	peopleCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	fmt.Println("")

	var result []string
	for i := 0; i < int(peopleCount); i++ {
		fmt.Print("Person composition : ")
		personComposition := strings.TrimSpace(readLine(reader))
		fmt.Println("")

		testReport := isPositive(virusComposition, personComposition)
		if testReport {
			result = append(result, "POSITIVE")

		} else {
			result = append(result, "NEGATIVE")

		}

	}
	for _, val := range result {
		fmt.Println(val)
	}
}
func isPositive(v string, b string) bool {
	NoOfCharfoundInVirusComposition := []bool{}
	for i := 0; i < len(b); i++ {
		subStrOfB := fmt.Sprintf("%c", b[i])
		result := strings.ContainsAny(v, subStrOfB)
		if result {
			if strings.Index(b, string(b[i])) <= strings.Index(v, string(b[i])) {
				NoOfCharfoundInVirusComposition = append(NoOfCharfoundInVirusComposition, result)
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

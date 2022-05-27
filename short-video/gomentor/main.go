package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Question List ")
	listofQuestion := questionList()
	for i := 1; i <= len(listofQuestion); i++ {
		fmt.Println(fmt.Sprint(i) + "." + listofQuestion[i])
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')

		qno, err := strconv.ParseInt(strings.Split(input, "\n")[0], 0, 0)
		if err != nil {
			err := fmt.Errorf("%v", err.Error())
			fmt.Println(err)
		}
		if qno == 3 {
			os.Exit(0)
		}
		fmt.Println("answer :- \n" + answerList(int(qno)))
	}

}

func questionList() map[int]string {
	questionList := map[int]string{
		1: "What is benefits of golang?",
		2: "Explain the use of interface in golang?",
		3: "Exit",
	}

	return questionList
}

func answerList(qno int) string {
	answerList := map[int]string{}

	answerList[1] = `1. Web Service: net/http
2. Database: database/sql
3. Container: container/list, container/heap
4. Compression: compress/gzip`

	return answerList[qno]
}

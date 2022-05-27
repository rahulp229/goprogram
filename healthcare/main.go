package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//// quiz app that read data from csv and each question must have 10 sec
	//and end result is no of correct answers and wrong answers show result
	f, err := os.Open("qna.csv")
	if err != nil {
		log.Fatalln("unableto read file ", err)
	}
	reader := csv.NewReader(f)
	data, _ := reader.ReadAll()

	//var questions []string
	//	answers := map[string]bool{}
	inputReader := bufio.NewReader(os.Stdin)
	noOfQuestionAnswerCorrect := 0
	noOfQuestionNotAnswer := 0

	for _, val := range data {

		answer := strings.Split(val[0], "=")[1]
		fmt.Println(strings.Split(val[0], "=")[0])
		//inputReader.Discard()
		//timeout := time.Durat
		in, errIn := inputReader.ReadString('\n')
		if errIn != nil {
			log.Fatalln("invalid input")
		}

		//time.Sleep(2 * time.Second)
		fmt.Println(">>", in)
		getAnswer := strings.Split(in, "\n")[0]
		if (getAnswer != "") && (getAnswer == answer) {
			noOfQuestionAnswerCorrect++
		} else {
			noOfQuestionNotAnswer++
		}

	}
	fmt.Println("correct = ", noOfQuestionAnswerCorrect)
	fmt.Println("incorrect = ", noOfQuestionNotAnswer)

	http.HandleFunc("/info1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	//fmt.Print(http.ListenAndServe(":8080", nil))

}

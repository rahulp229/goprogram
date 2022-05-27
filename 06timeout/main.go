package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	f, err := os.OpenFile("qna.csv", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("open file " + err.Error())

	}

	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic("unable to read file " + err.Error())
	}
	fmt.Println("data", data)
	//	var answer string = ""
	//	ch := make(chan string)
	result := make(chan string)
	var answer string

	for _, val := range data {
		go func() {

			select {
			case r := <-result:
				fmt.Println("answer is:", r)
			case <-time.After(10 * time.Second):
				fmt.Println("timeout")
				os.Exit(0)
			}

		}()
		fmt.Println(val)
		fmt.Scanln(&answer)
		result <- answer

	}

}

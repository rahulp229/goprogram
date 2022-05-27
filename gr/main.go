package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	file, _ := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	wrt := csv.NewWriter(file)
	var wg sync.WaitGroup
	//wg.Add(2)
	//fmt.Printf("waite grount :%v<<<",wg.Add(2))
	for i := 1; i <= 10; i++ {
		//wg.Add(1)
		wg.Add(2)
		go test(wrt, &wg, "first", i)
		//time.Sleep(time.Millisecond * 500)
		go test(wrt, &wg, "second", i)
		wg.Wait()
	}
	//	wg.Wait()

}

func test(w *csv.Writer, wg *sync.WaitGroup, str string, i int) {
	//defer
	defer func() {
		fmt.Println("done >>>", i)
		wg.Done()
	}()
	data := []string{fmt.Sprintf(" %v ==> %v \n", i, str)}
	errr := w.Write(data)
	if errr != nil {
		fmt.Println("error", errr)
	}
	fmt.Printf(" %v ==> %v \n", i, str)
	time.Sleep(time.Second)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Total = 10

func main() {

	var test interface{}
	test = "123"
	switch test.(type) {
	case int:
		fmt.Println("its int")
	case string:
		fmt.Println("its string")

	}
	fmt.Println("Hello, playground ", test)
	wg := sync.WaitGroup{}
	Addgr(wg, Total, 0, "")
	//metadata := []Dashboard{{Tabs: []Tab{{Rows:}{}}}}
	reader := bufio.NewReader(os.Stdin)
	action := map[string]bool{"add": true, "remove": true}
	for {
		fmt.Printf("enter the action : ")
		input, _ := reader.ReadString('\n')
		cmd := strings.Split(input, "\n")[0]

		fmt.Printf("no. of go routine %v \n ", cmd)
		input, _ = reader.ReadString('\n')
		ngr, _ := strconv.ParseInt(strings.Split(input, "\n")[0], 10, 64)

		if !action[cmd] {
			os.Exit(0)
		}

		if cmd == "add" {
			Total = Total + int(ngr)
			Addgr(wg, Total, int(ngr), cmd)
		}
		if cmd == "remove" {
			Total = Total - int(ngr)
			Addgr(wg, Total, int(ngr), cmd)
		}

	}

}

func Addgr(wg sync.WaitGroup, total int, stopCount int, action string) {
	wg.Add(total)
	command := make(chan bool, stopCount)
	for i := 1; i <= total; i++ {
		go Worker(&wg, i, command)
	}
	if action == "remove" {
		for i := 1; i <= stopCount; i++ {
			time.Sleep(250 * time.Millisecond)

			command <- true
		}
	}

	wg.Wait()
}

func Worker(wg *sync.WaitGroup, i int, ch chan bool) {
	defer wg.Done()
	select {
	case <-ch:
		return
	default:
		fmt.Printf("thread %v is working \n", i)

	}
}

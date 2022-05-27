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
	wg := sync.WaitGroup{}
	//total := 10
	//wg.Add(10)
	Addgr(Total, wg, "", 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("enter command : ")
		cmd, _ := reader.ReadString('\n')
		fmt.Println("command ", len(strings.Split(cmd, "\n")[0]))
		cmd = strings.Split(cmd, "\n")[0]
		cmdMap := map[string]bool{"insert": true, "remove": true, "stop": true}
		if !cmdMap[cmd] {
			os.Exit(0)
		}

		if cmd == "stop" {
			os.Exit(0)
		}

		fmt.Println("no.of go rountine want to ", cmd)
		input, _ := reader.ReadString('\n')
		fmt.Println("input ", input)
		grcount, err := strconv.ParseInt(strings.Split(input, "\n")[0], 10, 64)
		if err != nil {
			panic("error : " + err.Error())
		}

		//fmt.Println("innput ", grcount)
		//if Total < int(grcount) {
		if cmd == "insert" {

			updatedTotal := int(grcount) + Total
			Total = updatedTotal

			Addgr(int(updatedTotal), wg, cmd, int(grcount))

		}
		//if Total > int(grcount) {
		if cmd == "remove" {
			updatedTotal := Total - int(grcount)
			Total = updatedTotal
			Addgr(int(Total), wg, cmd, int(grcount))
		}

	}

}
func Removegr(total int, wg sync.WaitGroup, cmd string, stopCount int) {
}
func Addgr(total int, wg sync.WaitGroup, cmd string, stopCount int) {
	wg.Add(total)
	command := make(chan bool, stopCount)
	//
	for i := 1; i <= total; i++ {
		go worker(command, i, &wg)
	}
	if cmd == "remove" {
		for i := 0; i < stopCount; i++ {
			//fmt.Println("remove goroutine ", i)
			time.Sleep(250 * time.Millisecond)
			command <- true
		}

	}

	wg.Wait()

}

func worker(ch chan bool, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	//j := 0
	//	m := sync.Mutex{}
	//m.Lock()
	select {
	case <-ch:
		fmt.Println("stop gr")
		return
	default:
		fmt.Printf("thread is %v working \n", i)
	}
	//m.Unlock()
}

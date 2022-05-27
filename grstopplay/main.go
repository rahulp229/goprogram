package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	fmt.Println("inside routine")
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func main() {
	//	var wg sync.WaitGroup
	var wg1 sync.WaitGroup
	//	wg.Add(1)
	//	command := make(chan string)
	//	go routine(command, &wg)

	//time.Sleep(1 * time.Second)
	// time.Sleep(1500 * time.Millisecond)
	// command <- "Pause"

	// time.Sleep(1 * time.Second)
	// command <- "Play"

	// time.Sleep(5 * time.Second)
	// command <- "Stop"

	//	wg.Wait()

	wg1.Add(1)
	command1 := make(chan string)
	//	command1 <- "play"
	go showName(command1, &wg1)
	//command1 <- "pause"
	time.Sleep(1 * time.Second)
	command1 <- "pause"
	//command1 <- "play"
	time.Sleep(1 * time.Second)
	command1 <- "play"

	time.Sleep(1250 * time.Millisecond)
	command1 <- "stop"
	//command1 <- "pause"
	wg1.Wait()
}

func showName(chn chan string, wg *sync.WaitGroup) {
	fmt.Println("inside showName")

	defer wg.Done()

	//fmt.Println("cmd :", cmd)
	var status string = "play"
	//cmd := <-chn
	for {

		select {
		case cmd := <-chn:
			fmt.Println("cmd :", cmd)
			switch cmd {
			case "pause":
				status = "pause"

			case "play":
				status = "play"
			case "stop":
				return
			}
		default:
			if status == "play" {
				//fmt.Println("patil")
				time.Sleep(time.Millisecond * 250)
				fmt.Println("rahul")
			} else {
			}
		}
	}
	// for {

	// 	select {
	// 	case cmd := <-chn:
	// 		fmt.Println(cmd)
	// 		if cmd == "pause" {
	// 			//	time.Sleep(250 * time.Millisecond)
	// 			//fmt.Println("bbb")
	// 			status = "pause"
	// 		}

	// 		if cmd == "stop" {
	// 			return
	// 		}

	// 		if cmd == "play" {
	// 			status = "play"
	// 			//time.Sleep(250 * time.Millisecond)

	// 		}
	// 	default:
	// 		if status == "play" {
	// 			time.Sleep(time.Millisecond * 250)
	// 			fmt.Println("rahul")

	// 			//fmt.Println(i)
	// 		}
	// 	}

	// }
}

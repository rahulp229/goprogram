package main

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup
var stop bool = false

func printmsg(m chan string, o map[string]int) {
	defer Wg.Done()
	for {
		if stop == true && len(m) == 0 {
			return
		}
		c := <-m
		if _, ok := o[c]; ok {
			o[c] = o[c] + 1
		} else {
			o[c] = 1
		}
	}
}

func sendmsg(c chan string, s []string) {
	defer Wg.Done()
	for _, v := range s {
		c <- v
	}
	//i := <-c

	stop = true
	return
}

func main() {
	msg := make(chan string, 100)
	//ch := make(chan string)
	//ch <- "ping"
	//fmt.Println("ping")
	country := []string{"in", "us", "in", "in", "us"}
	out := make(map[string]int)
	go sendmsg(msg, country)
	Wg.Add(1)
	go printmsg(msg, out)
	Wg.Add(1)

	Wg.Wait()
	fmt.Println(out)
	fmt.Println("All done!")

	c := make(chan string)
	go count("hello",c)
}

func count (s string , ch chan string)  {
	ch <- s 
}
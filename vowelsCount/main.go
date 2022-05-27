// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

// problem statement -- //Find count of two consecutive vowels in any given string.
func main() {

	str := "aeqwsdieio"

	//vowels := []string{"a", "e", "i", "o", "u"}
	var result int
	tempmap := make(map[string]bool)
	tempmap["a"] = true
	tempmap["e"] = true
	tempmap["i"] = true
	tempmap["o"] = true
	tempmap["u"] = true
	counter := 0
	for i := 0; i < len(str); i++ {
		ch := string(str[i])
		//fmt.Println("---", tempmap[ch])
		fmt.Println("---", str[i])
		if tempmap[ch] {
			counter++
		}

		if counter == 2 {
			result++
			counter = 0
		}

	}

	fmt.Println("final count >>>", result)
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch := make(chan string)
	//ch <- "geeksforgeeks"
	//fmt.Println(<-ch)
	go A(&wg, ch)
	//fmt.Println(<-ch)
	go B(&wg, ch)
	fmt.Println(<-ch)
	go C(&wg, ch)
	fmt.Println(<-ch)
	//time.Sleep(time.Second)
	fmt.Println(<-ch)
	//ch <- "world"
	//
	//fmt.Println(<-ch)
	//ch <- "world"
	wg.Wait()

}

func A(wg *sync.WaitGroup, ch chan string) {
	fmt.Println("inside A")
	ch <- "hello B"
	//fmt.Println("A:", <-ch)
	//time.Sleep(time.Second)
	wg.Done()
	//fmt.Println("out", <-ch)
}

func B(wg *sync.WaitGroup, ch chan string) {
	fmt.Println("inside B")
	ch <- "hello A"
	//fmt.Println("B:", <-ch)
	//time.Sleep(time.Second)
	wg.Done()
	//fmt.Println("out", <-ch)
}

func C(wg *sync.WaitGroup, ch chan string) {
	fmt.Println("inside C")
	ch <- "hello AB"
	//fmt.Println("B:", <-ch)
	//time.Sleep(time.Second)
	wg.Done()
	//fmt.Println("out", <-ch)
}

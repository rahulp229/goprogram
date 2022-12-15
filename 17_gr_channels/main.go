package main

import (
	"fmt"
	"time"
)

func main() {
	numberCh := make(chan int)
	for i := 1; i <= 100; i++ {
		go func() {
			n := <-numberCh
			r := n * n
			fmt.Printf("square of %v is %v \n", n, r)
		}()
	}

	for i := 1; i <= 100; i++ {
		numberCh <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(numberCh)

}

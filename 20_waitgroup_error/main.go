package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

var noOftask int = 10
var grCount int = 2

func main() {

	job := make(chan int, noOftask)
	result := make(chan int, noOftask)

	for i := 1; i <= 2; i++ {
		//wg.Add(2)
		go worker(job, result, i)
	}

	for i := 1; i <= noOftask; i++ {
		job <- i
	}
	close(job)
	//wg.Wait()

	for i := 1; i <= noOftask; i++ {
		<-result
	}
	close(result)

	//
	counter := increment()
	fmt.Println("counter :", counter(2))
	fmt.Println("counter :", counter(2))

}

func increment() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func worker(jb <-chan int, res chan<- int, w int) {
	//defer wg.Done()
	//j := <-jb
	for job := range jb {
		fmt.Printf("worker %v started %v \n", w, job)
		time.Sleep(time.Millisecond * 1000)
		fmt.Printf("worker %v finished %v \n", w, job)
		res <- job
	}

}

//problem statement --> building work pool using goroutine whoose doing n no.of jobs
package main

import (
	"fmt"
	"time"
)

var NoOfWorker int = 3
var jobs = make(chan int, 5)
var result = make(chan int, 5)

func main() {
	for i := 1; i <= NoOfWorker; i++ {
		go worker(jobs, result, i)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for j := 1; j <= 5; j++ {
		// <-result

	}

	time.Sleep(time.Millisecond * 10000)
}

func worker(jobs <-chan int, result chan<- int, id int) {
	//temp := <-jobs
	//fmt.Println("in worker ", temp)
	for j := range jobs {

		fmt.Println("worker ", id, " started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, " finished job ", j)
		//result <- j * 2
	}
}

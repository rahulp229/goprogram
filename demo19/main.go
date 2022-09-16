package main

import (
	"fmt"
	"sync"
	"time"
)

var worker int = 3

func main() {
	wg := sync.WaitGroup{}
	wg.Add(worker)
	//m := sync.Mutex{}
	jobs := make(chan int, 10)
	result := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		//m.Lock()
		go job(&wg, jobs, i, result)
		//m.Unlock()
	}

	for j := 1; j < 10; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		<-result
	}
	wg.Wait()

}

func job(wg *sync.WaitGroup, job <-chan int, id int, results chan<- int) {
	defer wg.Done()
	//defer close(ch)
	//time.Sleep(time.Millisecond * 2000)
	for j := range job {
		fmt.Println("workder", id, " started job ", j)
		time.Sleep(time.Millisecond * 2000)
		fmt.Println("workder", id, " finished job ", j)
		results <- j * 2
	}

}

// emp (id, name , salary, dept, startDate, enddate)

// top 5 highest paid emp from engineering dept

// select * from Employee
// where salary = (select from Employee order by salary desc lime 5)
// where dept= " engineering" and
// where enddate = null

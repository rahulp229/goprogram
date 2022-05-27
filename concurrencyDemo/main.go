package main

import (
	"fmt"
	"sync"
)

var balance int = 1000

func main() {
	wg := sync.WaitGroup{}
	wg1 := sync.Mutex{}
	wg.Add(2)
	fmt.Println("banlace first : ", balance)
	go Withdrawn(700, &wg, &wg1)
	go Deposit(600, &wg, &wg1)
	wg.Wait()
	fmt.Println("banlace : ", balance)
}

func Withdrawn(amt int, wg *sync.WaitGroup, wg1 *sync.Mutex) {
	defer wg.Done()
	wg1.Lock()
	balance = balance - amt
	fmt.Println("banlace1 : ", balance)
	//return balance
	//	wg1.Unlock()
}

func Deposit(amt int, wg *sync.WaitGroup, wg1 *sync.Mutex) {
	defer wg.Done()
	wg1.Lock()
	balance = balance + amt
	fmt.Println("banlace2 : ", balance)
	//return balance
	wg1.Unlock()
}

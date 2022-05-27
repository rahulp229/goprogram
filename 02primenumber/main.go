package main

import (
	"fmt"
	"sync"
)

func main() {
	var primeNumbers []int
	var wg sync.WaitGroup
	//wg.Add(999999)
	var wgl sync.Mutex
	for i := 2; i <= 300; i++ {
		n := i
		wg.Add(1)
		go func(wg *sync.WaitGroup, wgl *sync.Mutex) {
			defer wg.Done()
			wgl.Lock()
			result := IsPrime(n)
			if result {
				primeNumbers = append(primeNumbers, n)
			}
			wgl.Unlock()
		}(&wg, &wgl)
		fmt.Println(">>>", primeNumbers)

	}
	wg.Wait()
	// for i := 2; i < 200000; i++ {
	// 	result := IsPrime(i)
	// 	if result {
	// 		primeNumbers = append(primeNumbers, i)
	// 	}
	// }
	fmt.Println("Prime number between 1 to 100 are : ", primeNumbers)
	fmt.Println("Prime number count between 1 to 100 are : ", len(primeNumbers))
}

func IsPrime(n int) bool {

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

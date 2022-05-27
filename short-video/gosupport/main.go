package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\nWhat is an interface?")
	str := `
An interface is a collection of method signatures 
that a Type can implement (using methods).
Hence interface defines the behavior of the object.
For example, 
a Man can walk and talk. If an interface defines method 
signatures for walk and talk while Man implements 
walk and talk methods, then Man is said to implement
that interface.
`
	//	fmt.Println(len(str))
	//	fmt.Println(string(str[48]))
	for i := 0; i < len(str); i++ {
		time.Sleep(35 * time.Millisecond)
		fmt.Printf("%v", string(str[i]))
	}
}

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

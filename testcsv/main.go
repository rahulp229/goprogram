package main

import (
	"fmt"
	"log"
	"testcsv/readFile"
)

func main() {
	csvFile, csvErr := readFile.OpenFile("1.csv")
	if csvErr != nil {
		log.Fatal("unable to open csv", csvErr)
		return
	}

	readErr := readFile.ReadFile(csvFile)
	if readErr != nil {
		log.Fatalf("unable to read csv %v", readErr)
	}
	fmt.Printf("%T>>", csvFile)
}

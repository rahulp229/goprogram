package main

import (
	"encoding/csv"
	"log"
	"os"
)

type User struct {
	Name  string
	Email string
}

func main() {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("unable to open file", err)
	}
	users := [][]string{{"john", "john@abc.com", "New york"}, {"adam", "adam@abc.com", "New jersy"}}
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	err = csvWriter.WriteAll(users)
	if err != nil {
		log.Fatal("unable to write the data", err)
	}

}

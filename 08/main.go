package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ponter struct {
	X int `json:"x"`
}

func main() {
	data := []byte(`{"x":1}`)
	//data := []byte{}
	var p Ponter
	if err := json.Unmarshal(data, &p); err != nil {

	} else {
		fmt.Println(p.X)
		fmt.Printf("type1 : %T", p)
		//	json.NewDecoder(p)
	}

	if b, err := json.Marshal(p); err != nil {
		fmt.Println()
	} else {
		fmt.Printf("type2 : %T", b)
	}

	// int to string
	j := 12
	i := strconv.Itoa(j)
	fmt.Printf("\ntype : %T", i)

	// string to int
	j1 := "12"
	i1, _ := strconv.Atoi(j1)
	fmt.Printf("\n type : %T ", i1)

//read input from terminal
	fmt.Println("\nread input from terminal")
	r := bufio.NewReader(os.Stdin)
//readstring
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("ReadString err : %v", err)
	}
	lengthOfInput := strings.Split(input, "\n")
	log.Println("ReadString input: ", input, len(input), len(lengthOfInput[0]))
//readLine
	fmt.Println("===read line===")
	line, isPrefix, err := r.ReadLine()
	if err != nil {
		log.Fatalln("ReadLine error : ", err)
	}
	log.Println("ReadLine line , isprefix, byte-string, len : ", line, isPrefix, string(line), len(string(line)))
//scanln
fmt.Println("===scanln===")
	scanInput, err := fmt.Scanln()
	if err != nil {
		log.Fatalln("scanln err: ", err)
	}
	log.Println("Scanln : ", scanInput)
// csv
	f, err := os.Open("names.csv")
	checkError(err, "csv")
	csvdata, _ := csv.NewReader(f).ReadAll()
	log.Println("csv : ", csvdata)

}
func checkError(e error, action string) {
	if e != nil {
		log.Fatalln("error")
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	var str string
	_ = json.Unmarshal(out, &str)
	fmt.Println(str)
	fmt.Printf("out: %v\n", str)
	fmt.Printf("err: %#v\n", err)
	if err != nil {
		log.Fatal(err)
	}
	//
	quit := make(chan bool)
	go func() {
		st := true
		for {
			select {
			case cme := <-quit:
				fmt.Println("cme :", cme)
				if cme == false {
					st = false
					//return
				}
			default:
				if st == true {
					time.Sleep(time.Millisecond * 250)
					fmt.Println("rr", st)
				}

			}
		}
	}()
	// …
	time.Sleep(time.Second * 1)
	quit <- true
	//time.Sleep(time.Second * 1250)
	quit <- false
	time.Sleep(time.Second * 1)

}

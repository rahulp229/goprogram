package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
- design an application where I can control number of go workers
Assume we start 10 go routines at the start of the application.  reads from command line new number of go routines and then stop/starts go routines as required.

Example:
  Application starts with 10 go routines.
  API call with 15
  App starts 5 new go routines
  API call with 3
  App stops 12 go routines

Database:
    Table: user_device_activity(user_email, device_serial_number, last_login_time)
    1. For each user find the most recent login time?
    2. For each user find the most recent login time along with the device on which the login was done?


1.	select last_login_time, Table,
	(select uid, max(last_login_time) as latestTime from Table group by uid) maxTime
	where Table.uid = maxTime.uid and
	Table.last_login_time = maxTime.lastestTime

2. 	select last_login_time, device_serial_number Table,
	(select uid, max(last_login_time) as latestTime from Table group by uid) maxTime
	where Table.uid = maxTime.uid and
	Table.last_login_time = maxTime.lastestTime
*/

var Total = 10

func main() {
	Addgr(Total, 0, "")
	reader := bufio.NewReader(os.Stdin)
	actionMap := map[string]bool{"add": true, "remove": true, "stop": true}
	for {
		fmt.Println("enter action :")
		action, _ := reader.ReadString('\n')
		cmd := strings.Split(action, "\n")[0]
		if cmd == "stop" {
			os.Exit(0)
		}
		fmt.Println("No.of go rountine  ")
		action2, _ := reader.ReadString('\n')
		cmd2 := strings.Split(action2, "\n")[0]
		noofgr, _ := strconv.ParseInt(cmd2, 10, 64)

		if !actionMap[cmd] {
			os.Exit(0)
		}

		if cmd == "add" {
			Total = Total + int(noofgr)
			Addgr(Total, int(noofgr), "")
		}

		if cmd == "remove" {
			Total = Total - int(noofgr)
			Addgr(Total, int(noofgr), cmd)
		}

	}
}

func Addgr(total int, stopCount int, action string) {
	wg := sync.WaitGroup{}
	wg.Add(total)
	command := make(chan bool, stopCount)
	for i := 1; i <= total; i++ {
		go Worker(&wg, i, command)
	}
	if action == "remove" {
		for i := 1; i <= stopCount; i++ {
			time.Sleep(250 * time.Millisecond)
			command <- true
		}
	}
	wg.Wait()
}

func Worker(wg *sync.WaitGroup, i int, ch chan bool) {
	defer wg.Done()
	select {
	case <-ch:
		return
	default:
		fmt.Println("i m gr ", i)
	}

}

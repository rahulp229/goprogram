// /*
// Input = [
//   {"id":1, "item": 2, "city": "Bengaluru"},
//   {"id":2, "item": 4, "city": "Mysore"},
//   {"id":1, "item": 5, "city": "Unnao"},
//   {"id":2, "item": 6, "city": "Agra"},
//   {"id":1, "item": 7, "city": "Kanpur"},
//   {"id":3, "item": 8, "city": "Ooty"}
// ]

// Output = [
//   {"id":3, "item": 8, "city": "Ooty"},
//   {"id":2, "item": 6, "city": "Agra"},
//   {"id":1, "item": 7, "city": "Kanpur"}
// ]
// zfw-jgee-abk
// */

// // type input struct {
// // 	id   int
// // 	item int
// // 	city string
// // }
// // type inputs []input

// // func main() {

// // 	in := inputs{{id: 1, item: 2, city: "mumbai"}, {id: 2, item: 5, city: "delhi"}, {id: 2, item: 6, city: "chenai"}}
// // 	fmt.Println("input :-> ", in)

// // 	temp := map[int]input{}
// // 	//result := inputs{}
// // 	for _, value := range in {
// // 		if temp[value.id].id != 0 {
// // 			item := input{id: value.id, item: value.item, city: value.city}
// // 			temp[value.id] = item

// // 		} else {

// // 		}
// // 	}
// // 	//append(result)
// // 	fmt.Println("reslut >> ", temp)

// // }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )

// type ReleasesInfo struct {
// 	Id      uint   `json:"id"`
// 	TagName string `json:"tag_name"`
// }

// // Function to actually query the GitHub API for the release information.
// func getLatestReleaseTag(repo string) (string, error) {
// 	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
// 	response, err := http.Get(apiUrl)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	releases := []ReleasesInfo{}

// 	if err := json.Unmarshal(body, &releases); err != nil {
// 		return "", err
// 	}

// 	tag := releases[0].TagName

// 	return tag, nil
// }

// // Function to get the message to display to the end user.
// func getReleaseTagMessage(repo string) (string, error) {
// 	tag, err := getLatestReleaseTag(repo)
// 	if err != nil {
// 		return "", fmt.Errorf("Error querying GitHub API: %s", err)
// 	}

// 	return fmt.Sprintf("The latest release is %s", tag), nil
// }

// func main() {
// 	msg, err := getReleaseTagMessage("docker/machine")
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, msg)
// 	}

// 	fmt.Println(msg)
// }

// Write 3 go routine function
// Function1 print numbers from 1 to N divisible by 3
// Function2 print numbers from 1 to N divisible by 5
// Function3 print numbers from 1 to N divisible by 15
package main

import (
	"fmt"
)

func main() {
	// (func(a aa) A)()
	// (func(b bb) A)()
	//wg := sync.WaitGroup{}
	//wg.Add(3)
	ch := make(chan int, 3)

	go function1(ch)
	//time.Sleep(time.Second)
	go function2(ch)
	go function3(ch)

	ch <- 100
	//time.Sleep(time.Second)
	ch <- 50
	//time.Sleep(time.Second)

	ch <- 150
	//.Sleep(time.Second)

	//close(ch)

}

func function1(n <-chan int) {

	num := <-n
	for i := 1; i <= num; i++ {
		select {
		case num := <-n:
			
		}
		if i%3 == 0 {
			fmt.Println(i, "devisible by 3 ")
		}
	}
	//time
}

func function2(n chan int) {
	num := <-n
	for i := 1; i <= num; i++ {
		if i%5 == 0 {
			fmt.Println(i, "devisible by 5 ")
		}
	}
}

func function3(n chan int) {
	num := <-n
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			fmt.Println(i, "devisible by 15 ")
		}
	}
}

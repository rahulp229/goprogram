package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {

	f, err := os.OpenFile("config.json", os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic(err.Error())
	}
	data, errd := os.ReadFile(f.Name())
	if errd != nil {
		panic(errd.Error())
	}

	dd := fmt.Sprintf("%s", data)
	fmt.Println(dd)
	data1, err1 := ioutil.ReadFile(f.Name())
	if err1 != nil {
		panic(err1.Error())
	}

	var s struct{ Name string }
	s.Name = "rr"
	fmt.Println(s)
	if err := json.Unmarshal(data1, &s); err != nil {
		panic(err.Error())
	}
	fmt.Println(s)

	//GET
	//var p []byte
	resp, _ := http.Get("http://localhost:9090/users")
	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("first", string(r))
	client := http.Client{}

	req := http.Request{Method: "GET", URL: &url.URL{Host: "http://localhost:9090", Path: "/users"}, Proto: "HTTP/1.0", }
	resp2, e1 := client.Do(&req)
	if e1 != nil {
		fmt.Println("error1:  ", e1)
	}
	r2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println("sencod", string(r2))

}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const Apikey = "AIzaSyAiSbN_p3W_6mK-SDCOcvaRPk4bSr6LKq8"

func main() {

	res, err := http.Get("https://abusiveexperiencereport.googleapis.com/v1/violatingSites?key=" + Apikey)
	if err != nil {
		log.Printf("get request failed %v ", err)
		os.Exit(1)
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("unable to fetch data : ", err)
	}
	fmt.Println("response : ", string(resData))
}

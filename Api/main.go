package main

import (
	"Api/router"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() { 
	router := router.Init()
	//	u := User{}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const connectString = "rahulp229:$wKv&TP_Y(}K@tcp(166.62.28.117:3306)/rkp19"

type DBConfig struct {
	Host         string
	Port         string
	UserName     string
	Password     string
	DatabaseName string
}

func main() {

	// create a database object which can be used
	// to connect with database.
	//user := "rahulp229"
	//password := "$wKv&TP_Y(}K"
	dbConfig := DBConfig{
		Host:         "166.62.28.117",
		Port:         "3306",
		UserName:     "rahulp229",
		Password:     "",
		DatabaseName: "",
	}
	connectionString := 
	db, err := sql.Open("mysql", "rahulp229:$wKv&TP_Y(}K@tcp(166.62.28.117:3306)/rkp19")

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	// Now its time to connect with oru database,
	// database object has a method Ping.
	// Ping returns error, if unable connect to database.
	err = db.Ping()

	// handle error
	if err != nil {
		panic(err)
	}

	fmt.Print("Pong\n")
	row := db.QueryRow("select * from test")
	// database object has a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	result := struct {
		Id     int32  `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Mobile string `json:"mob"`
	}{}
	if err = row.Scan(&result.Id, &result.Name, &result.Email, &result.Mobile); err != nil {
		log.Fatal("error : ", err)
	}
	fmt.Println("result : ", result)
	defer db.Close()

}

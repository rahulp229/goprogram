package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fileName := "sample.jpg"
	//	URL := "http://www.africau.edu/images/default/sample.pdf"
	URL := "https://storage.googleapis.com/everledger-platform-uat/b85dcada-c414-477c-ba4e-e4173a12df24/1d6aff57-159f-4e2a-99ec-88ebfbbcf00c/8fd2de37-5664-47df-9fbc-8f4bf6ccfe71/pink-neon.jpg?Expires=1639489591&GoogleAccessId=everledger-uat-bucket%40everledger-internal.iam.gserviceaccount.com&Signature=ng%2B%2B%2FhVSigHqNYyp0DnjbSUC%2BiYeHF1X4AJTjMAQn1A9Q0P1sEqSGUYt92wpMNSvnvVXmJJVGQ5mCpUnXq9UJWFK9b0RKj8l%2BqWVZmbkWXsfrojQ9rKNSVpfXqh%2B5nWCFnesyWfd186hi8PX09i6YuTyfv%2BW8InrzH68Ma4ogSOqMSuKnjtLlPSHC1MQDkqIzF8%2F2zLv4wCuUQs87oI73ahIkOtzPBQmgSM4SPfaNgI9Bn5sPDkxG7Jzs9sYtV7bY1tNEt3CI64YtDDjwAIfX6S%2FLVa0sJPfxFqOxsDD9XKZxPAmQsbUnknfnACVKAw8Uh4dmM7vSAjw%2BYrcDpj%2BPA%3D%3D"
	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %s downlaod in current working directory", fileName)
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Println(response.StatusCode)
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	uu, ee := url.Parse("https://storage.googleapis.com/everledger-platform-uat/b85dcada-c414-477c-ba4e-e4173a12df24/1d6aff57-159f-4e2a-99ec-88ebfbbcf00c/8fd2de37-5664-47df-9fbc-8f4bf6ccfe71/pink-neon.jpg?Expires=1639162635&GoogleAccessId=everledger-uat-bucket%40everledger-internal.iam.gserviceaccount.com&Signature=CEEEf0Ydu6u2lzvDKKlF%2FRKepqKB1Z6YqJgW8yca5MSi9lBEGKuyHw0lqFVmRe9VpZvKvjK43IydhXa0q5TogtwCQYUCpHrzH%2B2WUK7bQmDhS%2FesBaMEhQOJdyU7FFeSKl4QAZOLcpaOP%2FnAAdS%2ByABw9MTuZtxZ68ozowVyOjSZWDowTYrEOQyh9v05gXJDCox4Fw9N3JWKBOy0lQO3pX1%2BEzrgPzu9KJmDzZ2EeG3oSUVrI2s4hlFW0JJgWdx1GwaDiFFxD7ai7eqMDGWZdkYyf2mXdCXt66I0mDoiyP80CEvPl9VJQOLi8ZHDK419ywXjxGDhKTDC2nr2AXgeWQ%3D%3D")
	fmt.Println("errrooorr>>>> ", ee)
	fmt.Println("uuuurrrlll>>>> ", uu)

	f, _ := os.Open("sample.jpg")
	stats, _ := f.Stat()
	fmt.Println("extraInfo1 ", stats.Sys())

	return nil
}

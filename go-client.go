package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//make a sample HTTP GET Request
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	//check for nil response
	if err != nil {
		log.Fatal(err)
	}

	//read all response body
	data, _ := ioutil.ReadAll(res.Body)

	//close the response body
	res.Body.Close()

	//print data as string
	fmt.Printf("%s\n", data)

	//make a sample HTTP GET Request to go-webserver
	res2, err2 := http.Get("http://localhost:9090/search")

	//check for nil response
	if err2 != nil {
		log.Fatal(err2)
	}

	//read all response body
	data2, _ := ioutil.ReadAll(res2.Body)

	//close the response body
	res.Body.Close()

	//print data as string
	fmt.Printf("%s\n", data2)

}

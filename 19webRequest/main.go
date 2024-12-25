package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://developers.lco.global/"

func main() {
	fmt.Println("Web Request in go-lang")
	response, err := http.Get(url)

	checkError(err)

	fmt.Printf("Resposse is of type: %T\n", response)
	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	checkError(err)
	content := string(databyte)
	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

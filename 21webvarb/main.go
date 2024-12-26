package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Wellcome in fo WebVarb - LCO")
	performGetRequest()
}

func performGetRequest() {
	const myUrls = "http://localhost:8000/get"

	response, err := http.Get(myUrls)
	checkError(err)
	defer response.Body.Close()
	
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseSting strings.Builder
	content , err := ioutil.ReadAll(response.Body)
	checkError(err)

	byteCount, _ := responseSting.Write(content)

	fmt.Printf("Byte count is: %v\n", byteCount)
	fmt.Println(responseSting.String())
	
	// fmt.Println(string(content))
	
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error ", err)
	}
}

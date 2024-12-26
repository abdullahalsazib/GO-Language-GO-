package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const urls = "https://developers.lco.global/"

func main() {
	basicUrls()
	// Parsing url
	result, err := url.Parse(urls)
	checkError(err)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)

}

func basicUrls() {
	fmt.Println("Wellcome to go - lang url's")

	fmt.Println(urls)

	response, err := http.Get(urls)
	checkError(err)
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	// databyte, err := ioutil.ReadAll(response.Body)
	// checkError(err)
	// content := string(databyte)
	// fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error ", err)
	}
}

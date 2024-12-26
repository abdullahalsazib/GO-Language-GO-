package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Wellcome in fo WebVarb - LCO")
	// PerformGetRequest()
	// PerfromPostRequest()
	PrefromPostFormRequest()
}

func PerformGetRequest() {
	const myUrls = "http://localhost:8000/get"

	response, err := http.Get(myUrls)
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseSting strings.Builder
	content, err := ioutil.ReadAll(response.Body)
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

// Post request
func PerfromPostRequest() {
	const myUrls = "http://localhost:8000/post"

	//fake json payload
	requestBoey := strings.NewReader(`
		{
			"coursename": "let's go with golang",
			"price": 0,
			"Platform": "Udemy"
		}
	`)

	response, err := http.Post(myUrls, "application/json", requestBoey)
	checkError(err)
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))


}


// Post From
func PrefromPostFormRequest() {
	const myUrls = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Jack")
	data.Add("lastname", "Sparrow")
	data.Add("email", "jackemail@mail.com")
	
	response, err := http.PostForm(myUrls, data)
	checkError(err)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
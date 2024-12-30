package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var myUrl = "https://jsonplaceholder.typicode.com/todos"

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD....")
	PostRequest()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetMethod() {
	fmt.Println("GET - Method")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	checkError(err)

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response: ", res.StatusCode)
	}

	// data, err := ioutil.ReadAll(res.Body)
	// checkError(err)
	// defer res.Body.Close()
	// fmt.Println(string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	checkError(err)

	fmt.Println(todo.UserID)
}

func PostRequest() {
	fmt.Println("POST - Method")
	todo := Todo{
		UserID:    213,
		Title:     "just lorem text",
		Completed: false,
	}

	// convert into json

	jsonData, err := json.Marshal(todo)
	checkError(err)

	// convert json data to string !

	jsonString := string(jsonData)

	// convert string to json reader

	jsonReader := strings.NewReader(jsonString)

	// http post method
	res, err := http.Post(myUrl, "aplication/json", jsonReader)
	checkError(err)
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(data))

}

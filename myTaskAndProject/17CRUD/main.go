package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD....")
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("welcome in url")
	u := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("Type is %T\n", u)

	ur, err := url.Parse(u)
	if err != nil {
		panic(err)
		return
	}
	
	fmt.Printf("ur type %T\n", ur)
	fmt.Println(ur)
}

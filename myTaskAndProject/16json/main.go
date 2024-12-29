package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Fullname string `json:"name"`
	Age      int    `json:"age"`
	IsAdult  bool   `json:"is_Adult"`
}

func main() {
	fmt.Println("Welcome on go - json <|")
	person := Person{Fullname: "jack", Age: 45, IsAdult: true}

	// encoding or marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(string(jsonData))

	// deconding or unmarshal

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		panic(err)
	}
	fmt.Println(personData.Fullname)
	fmt.Println(personData.Age)
	fmt.Println(personData.IsAdult)

}

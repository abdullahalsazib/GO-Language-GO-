package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
    Name string
}



func main() {
    fmt.Println("Map in JSON")

    a := make(map[string]employee)
    a["1"] = employee{Name: "John"}
    a["2"] = employee{Name: "Jack"}

    j, err := json.Marshal(a)

    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
    } else {
        fmt.Println(string(j))
    }
    fmt.Println("---------------")
    
    var b map[int]employee
    json.Unmarshal(j, &b)

    fmt.Println(b)

    keyDelete()
}

// Map key delete
func keyDelete() {
    smaple := make(map[string]int)

    smaple["a"] = 1
    smaple["b"] = 2
    smaple["c"] = 3
    smaple["d"] = 4

    fmt.Println(smaple)

    // Check and delete
    if _, ok := smaple["a"]; ok {
        delete(smaple, "a")
    }
    fmt.Println(smaple)
    // Directly delte 
    delete(smaple, "c")
    fmt.Println(smaple)
}

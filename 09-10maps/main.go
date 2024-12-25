package main

import (
	"encoding/json"
	"fmt"
)

func main() {
    fmt.Println("Maps in go")
    language := make(map[string]string)
    language["JS"] = "javaScript"
    language["JAVA"] = "JAVA"
    language["PY"] = "Python"

    fmt.Println("List of all Language's ", language)
    fmt.Println(language["JS"])

    delete(language, "PY")
    fmt.Println(language)
    
    // loops are interasting in golan

    for key, value := range language{
        fmt.Printf("For kay %v value %v\n", key, value)
    }

    sample := map[string]string {
        "a": "x",
        "b": "y",
    }
    fmt.Println(sample)

    for _, value := range sample {
        fmt.Println(value)
    }

    keys := getAllKeys(sample)

    fmt.Println(keys)

    fmt.Printf("Length of sample: %v\n", len(sample))


    // Declear 
    employeeSalary := make(map[string]int)

    // Add a key value
    employeeSalary["Tom"] = 2000
    employeeSalary["sam"] = 1200

    lenOfMap := len(employeeSalary)
    fmt.Println(lenOfMap)

    // Map to JSON 
    a := make(map[int]string)
    a[1] = "John"

    j, err := json.Marshal(a)
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    } else {
        fmt.Println(string(j))
    }
}

func getAllKeys(smaple map[string]string) [] string {
    var keys []string
    for k := range smaple {
        keys = append(keys, k)
    }
    return keys
}

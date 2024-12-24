package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Welcome to time study of golang..")
    persentTime := time.Now()
    //fmt.Print(persentTime)

    fmt.Println(persentTime.Format("01-02-2006 | 15:04:03 | Monday"))

    createdDate := time.Date(2020,time.August, 10, 23, 23,0, 0, time.UTC)
    fmt.Println(createdDate)
    fmt.Println(createdDate.Format("01-02-2006 wednesday"))
}

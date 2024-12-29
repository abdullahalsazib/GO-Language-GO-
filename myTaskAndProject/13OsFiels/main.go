package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
    fmt.Println("Wellcome in os - File's")
    /*
    // Create a file 
    file, err := os.Create("./Example.txt")
    if err != nil {
        fmt.Println("File Createing some issu:", err)
        return
    }
    // Close the file
    defer file.Close()

    content := "Hello, World"
    _, errore := io.WriteString(file, content+"\n")
    if errore != nil {
        panic(errore)
    }
    fmt.Println("file create successfully....")
    */ 
    /*
    //open file 
    file, err := os.OpenFile("Example.txt")
    if err != nil {
        panic(err)
        return
    }
    defer file.Close()

    // buffer create 
    buffer := make([]byte, 1024)

    for {
        file, err := file.Read()
        if err == io.EOF() {
            break
        }
        if err != nil {
            panic(err)
        }
    }
    */









}

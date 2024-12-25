package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Work with -> file's OS")
	content := "this need s to go in a file - mdsazibhossin2021@gmial.com"

	file, err := os.Create("./myLcogofile.txt")

	if err != nil {
		panic(err)
	}
	lenght, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", lenght)
	file.Close()

	readFile("./myLcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	// fmt.Println("Text data inside the file is:\n", databyte) // Byte Code Return <<=
	fmt.Println("Text data inside the file is:\n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

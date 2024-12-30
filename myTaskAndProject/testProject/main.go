package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

var c = color.New(color.FgGreen).Add(color.Underline)
var red = color.New(color.FgHiRed).Add(color.Underline)

/*

{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}

*/

type Person struct {
	UserID    string `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Hello, World")

	/*
		    r := mux.NewRouter()
			r.HandleFunc("/", homeServer)

			// Create a server |> GO =>
			fmt.Println("Server running on http://localhost:8080")
			http.ListenAndServe(":8080", r)
	*/

	r := mux.NewRouter()
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	

	res, err := http.Get(myUrl)
	checkError(err)
	defer res.Body.Close()
	checkStatusCode(res.StatusCode)
	fmt.Printf("%T\n", res.StatusCode)

	var mainDAta *Person

	mainDAta = JsonEncoding(res)
	fmt.Println(mainDAta.Title)


    creatStaticServer(r)

}

func JsonEncoding(res *http.Response) *Person {
	byteData, err := ioutil.ReadAll(res.Body)
	checkError(err)

	// json unmarsele
	var data Person
	json.Unmarshal(byteData, &data)

	return &data
}

func creatStaticServer(r *mux.Router) {
	staticDir := "./static"

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))

	// Starting server
	fmt.Println("Server running .....")
	fmt.Println("Server running .....")
	c.Println("http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

// func homeServer(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello, World..!</h1>"))
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkStatusCode(code int) {
	if code != http.StatusOK {
		red.Println("Request Error", code)
	}
}

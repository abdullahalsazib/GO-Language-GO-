package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Wellcome |> to JSON")
	// EncodeJson()
	DecodeJson()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodeJson() {
	locCourses := []course{
		{"ReactJs", 299, "udemy", "1234", []string{"webdev", "js"}},
		{"AngularJs", 239, "udemy", "332laxe", []string{"webdev", "js"}},
		{"CEH_v9", 399, "udemy", "22aadd", nil},
	}

	finalJson, err := json.MarshalIndent(locCourses, "", "\t")
	checkError(err)

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJs",
                "Price": 299,
                "website": "udemy.com",
                "tags": [
                        "webdev",
                        "js"
                ]
        }
	`)
	// <^> <| <--> 

	 var locCourses course

	 checkValue := json.Valid(jsonDataFromWeb)
	 if checkValue {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &locCourses)
		fmt.Println("%#v\n", locCourses)

	 } else {
		fmt.Println("JSON was not valid!")
	 }

	 // some cases where you just want to add data to key value

	 var myOnlneData map[string]interface{}
	 json.Unmarshal(jsonDataFromWeb, &myOnlneData)

	 fmt.Printf("%#v", myOnlneData)

	 for k, v := range myOnlneData {
		fmt.Printf("key is %v and value is %v and type of %T\n", k, v, v)
	 }

}

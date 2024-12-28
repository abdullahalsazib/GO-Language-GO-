package main

import "fmt"

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB - Database
var courseDB []Course

// Middleware, helper - file
func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("welcome in build api")

}

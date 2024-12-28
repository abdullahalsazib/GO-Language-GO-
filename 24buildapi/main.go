package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json"courseprice"`
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
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("welcome in build api")

	r := mux.NewRouter()

	// seeding
	courseDB = append(courseDB, Course{
		CourseId:    "1",
		CourseName:  "Go",
		CoursePrice: 100,
		Author: &Author{
			Fullname: "Jack Sparrow",
			Website:  "https://learncodeonline.in",
		},
	})

	courseDB = append(courseDB, Course{
		CourseId:    "2",
		CourseName:  "Node",
		CoursePrice: 200,
		Author: &Author{
			Fullname: "John Doe",
			Website:  "https://learncodeonline.in",
		},
	})

	courseDB = append(courseDB, Course{
		CourseId:    "3",
		CourseName:  "React",
		CoursePrice: 300,
		Author: &Author{
			Fullname: "Sparrow Doe",
			Website:  "https://learncodeonline.in",
		},
	})

	// routhing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOnecourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOnecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOnecourse).Methods("DELETE")

	// listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))

}

// controller file

// serve home page
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Go home page or Api by Jack Sparrow</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Corses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courseDB)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Corses")
	w.Header().Set("Content-Type", "application/json")

	// grap id from request
	params := mux.Vars(r)

	// loop through course, fiind mathicing id and return the response
	for _, course := range courseDB {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with this id")
	return
}

func createOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Corses")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if IsEmpty(&course) {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	for _, c := range courseDB {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}

	// generate unique id, string
	// append to courseDB

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courseDB = append(courseDB, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Corses")
	w.Header().Set("Content-Type", "application/json")

	// first - grap id from requesr
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courseDB {
		if course.CourseId == params["id"] {
			courseDB = append(courseDB[:index], courseDB[:index+1]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courseDB = append(courseDB, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: sned a response when id is not found
	for _, course := range courseDB {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode("No course found with this id")
		}
	}

}

func deleteOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Corses")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop, id, remove (index, index+1)
	for index, course := range courseDB {
		if course.CourseId == params["id"] {
			courseDB = append(courseDB[:index], courseDB[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(params["id"] + " deleted")
}

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

// Create a Author
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courseDB []Course

// Middleware, hepler - file
func IsEmpty(c *Course) bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Hello, World!")
	r := mux.NewRouter()

	// Seeding
	courseDB = append(courseDB, Course{
		CourseId:    "1",
		CourseName:  "Go",
		CoursePrice: 100,
		Author: &Author{
			FullName: "Jack",
			Website:  "https://udemy.com",
		},
	})
	courseDB = append(courseDB, Course{
		CourseId:    "2",
		CourseName:  "Java",
		CoursePrice: 160,
		Author: &Author{
			FullName: "Sparrow",
			Website:  "https://udemy.com",
		},
	})

	// routing -->
	r.HandleFunc("/", homeServer).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Define a custom 404 handler
	r.NotFoundHandler = http.HandlerFunc(curstom404Handler)

	// Surve a localhost server http://localhost:8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func homeServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>HOME Server, Hello, World!</h1>"))
}
func curstom404Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<div style="
	width: 100%;
    height: 100vh;
    background: gray;
    margin: 0;
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: center;
	text-align: center;
	">
	<h1>404 PAGE !</h1>
	<button onclick="alert('hello')">click me</button>
	</div>`))
}

// Check Error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get the all Courses!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courseDB)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grap id from request
	parems := mux.Vars(r)

	// loop thorugh course, find mathiching id and return the response
	for _, course := range courseDB {
		if course.CourseId == parems["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No course found with this id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data..!")
	}

	// What about - {}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	checkError(err)

	if IsEmpty(&course) {
		json.NewEncoder(w).Encode("No Data Inside JSON..!")
	}

	for _, c := range courseDB {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course alrady exists..!")
		}
	}

	// // append to courseDB

	rand.Seed(time.Now().UnixMicro())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courseDB = append(courseDB, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grap id form request
	params := mux.Vars(r)

	// loop, id , remove, add with my ID
	for index, course := range courseDB {
		if course.CourseId == params["id"] {
			courseDB = append(courseDB[:index], courseDB[:index+1]...)
			var course Course
			err := json.NewDecoder(r.Body).Decode(&course)
			checkError(err)
			course.CourseId = params["id"]
			courseDB = append(courseDB, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: send a response when id is not found
	for _, course := range courseDB {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode("No Course found with this id")
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop, id, remove,(index, index+1)

	for index, course := range courseDB {
		if course.CourseId == params["id"] {
			courseDB = append(courseDB[:index], courseDB[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(params["id"] + "Deleted..!")
}

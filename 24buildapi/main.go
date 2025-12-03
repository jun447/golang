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
	CourseId   string   `json:"courseid"`
	CourseName string   `json:"coursename"`
	Price      int      `json:"price"`
	Author     *Author  `json:"author"`
	Tags       []string `json:"tags"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}
type ErrorResponse struct {
	Error string `json:"error"`
	ID    string `json:"id,omitempty"`
}

//db simulation slice

var courses []Course

//simulating middleware hlper functions

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Hello world")
	//seeding the course db
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", Price: 299, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}, Tags: []string{"frontend", "javascript"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "MERN Stack", Price: 499, Author: &Author{Fullname: "Jane Smith", Website: "janesmith.com"}, Tags: []string{"fullstack", "javascript"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Golang", Price: 399, Author: &Author{Fullname: "Alice Johnson", Website: "alicejohnson.com"}, Tags: []string{"backend", "golang"}})
	
	//routing
	r:=mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",DeleteOneCourse).Methods("DELETE")




	//listen to port 4000
	// var r *mux.Router
	log.Fatal(http.ListenAndServe(":4000",r))
}

// controllers -
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome Learning Bro</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder((w)).Encode(courses)
}
func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Courses y Id")
	w.Header().Set("Content-Type", "application/json")
	//grabing id
	params := mux.Vars(r)
	fmt.Println(params)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder((w)).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id " + params["id"])
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Courses y Id")
	w.Header().Set("Content-Type", "application/json")

	//emptybody
	if r.Body == nil {
		json.NewEncoder(w).Encode("Insan Ban")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Insan Ban")

	}

	for _,value :=range courses{
		if value.CourseName==course.CourseName{
			json.NewEncoder(w).Encode("alredy exits bro")
		}
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa( rand.Intn(100))



	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Courses by Id")
	w.Header().Set("Content-Type", "application/json")
	//grabing id
	params := mux.Vars(r)
	fmt.Println(params)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// json.NewEncoder((w)).Encode(course)
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted Successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id " + params["id"])
}
func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Courses by Id")
	w.Header().Set("Content-Type", "application/json")
	//grabing id
	params := mux.Vars(r)
	fmt.Println(params)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// json.NewEncoder((w)).Encode(course)
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses, course)
			json.NewEncoder(w).Encode("Updated Successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id " + params["id"])
}

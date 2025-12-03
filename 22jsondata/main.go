package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hai json")

	// encodejson()
	decodejson()

}

// write above code as func named encodejson
func encodejson() {
	ourCourses := []course{
		{"ReactJS Bootcamp", 2999, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 1999, "Udemy", "def456", []string{"web-dev", "js"}},
		{"VueJS Bootcamp", 2999, "Udemy", "ghi789", nil},
	}

	// finaljson,_:= json.Marshal(ourCourses)
	finaljson, _ := json.MarshalIndent(ourCourses, "", " ")
	//  fmt.Println(string(finaljson))
	fmt.Printf("%s\n", finaljson)
}

func decodejson() {
	jsonDataFromWeb := []byte(`

	{
  "coursename": "ReactJS Bootcamp",
  "Price": 2999,
  "Platform": "Udemy",
  "Tags": [
   "web-dev",
   "js"
  ]
 }
	
	`)
	var course course
	isdataValid := json.Valid(jsonDataFromWeb)

	fmt.Println(isdataValid)
	if isdataValid {

		err := json.Unmarshal(jsonDataFromWeb, &course)
		handleNil(err)

		// fmt.Println(course)
	} else {
		fmt.Println("The json is not valid bro")
	}
	// poping dta into an map key,value
	var courseMap map[string]interface{}
	err:=json.Unmarshal(jsonDataFromWeb,&courseMap)
	handleNil(err)
	fmt.Println(courseMap)
	fmt.Printf("%#v\n",courseMap)

}
func handleNil(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

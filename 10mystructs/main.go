package main

import "fmt"

func main() {

	junaid := User{
		Name:   "Junaid",
		Email:  "junaid@example.com",
		Status: true,
		Age:    30,
	}
	fmt.Println(junaid)
	fmt.Printf("this person details are given below %+v",junaid)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

package main

import (
	"fmt"
	"time"
)
func main() {
	presentTime := time.Now()
	// fmt.Println("Welcome to MyTime package to manage time!")
	// fmt.Println("Current time is", presentTime)
	// // fmt.Printf("Type of presentTime is %T \n", presentTime)
	
	// // format time
	// fmt.Println("Formatted time is ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// // create date
	// createdDate:= time.Date(2020,time.November,10,23,0,0,0,time.UTC)

	// fmt.Println("Created date is ", createdDate.Format("01-02-2006 15:04:05 Monday"))

	//exploring nanoseconds
	fmt.Println("Nanoseconds are ", presentTime.UnixNano())
}
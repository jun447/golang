package main

import (
	"fmt"
	"net/url"
	"reflect"
)

// writing url with parameters and handling nil error
const myurl string="https://lco.dev:30000/learn?coursename=reactjs&paymentid=12345"

func main(){
	fmt.Println("welcome to urls in golang")
	// fmt.Println("url is ", myurl)
	// /parsing the url
	result,err:=url.Parse(myurl)
	handleNil(err)
	// fmt.Println("Parsed URL        :", result)
	// fmt.Println("Scheme            :", result.Scheme)
	// fmt.Println("Host              :", result.Host)
	// fmt.Println("Path              :", result.Path)
	// fmt.Println("Port              :", result.Port())
	// fmt.Println("Raw Query         :", result.RawQuery)

// 	// parse query params for nicer output
	queryParams := result.Query()
	fmt.Println("Query Parameters  :", queryParams)
		fmt.Println(reflect.TypeOf(queryParams).Kind())  

	fmt.Println("Course Name       :", queryParams["coursename"])
// 	fmt.Println("Payment ID        :", queryParams.Get("paymentid"))

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "devjun.netlify.app",
		Path: "projects",
	}
	fmt.Println(partsOfUrl)
}

func handleNil(e error){
	if e!=nil{
		fmt.Println(e)
		panic(e)
	}
}
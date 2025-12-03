package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	// "strings"
	"reflect"
)

func main() {
	fmt.Println("webrequests with json")
	// PerformeGetRequest("http://localhost:8000/get")
	// PerformPostJsonRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")
}

func PerformeGetRequest(myurl string) {

	// response,err := http.Get(myurl)
	var response *http.Response
	var err error
	response, err = http.Get(myurl)
	handleNil(err)

	// fmt.Println("response is ",*response)
	fmt.Println("response is ", &response)
	fmt.Println(reflect.TypeOf(response).Kind())
	fmt.Println(reflect.TypeOf(*response).Kind())

	contentByte, err := io.ReadAll(response.Body)
	handleNil(err)

	fmt.Println(string(contentByte))
	// var responseString strings.Builder

	// byteCount,err := responseString.Write(contentByte)
	// handleNil(err)

	// fmt.Println("byteCount is ",byteCount)

	// actualContent:= responseString.String()

	// fmt.Println("Ths is actual content by 2nd Method",actualContent)
	defer response.Body.Close()

	fmt.Println("StatusCode is ", response.StatusCode)
	fmt.Println("ContentLength is ", response.ContentLength)
}

func PerformPostJsonRequest(myurl string) {
	// our json data
	requestBody := strings.NewReader(`
		{
			"name":"True Crime Season 1",
			"price":"10000",
			"Platform":"HBO"
		}
	
	`)
	response,err:=http.Post(myurl,"application/json",requestBody)
	handleNil(err)
	fmt.Println("response from server is",response)
	fmt.Println("response body from server is",response.Body)

	contentByte, err := io.ReadAll(response.Body)
	handleNil(err)

	fmt.Println(string(contentByte))

	defer response.Body.Close()
}
func PerformPostFormRequest(myurl string) {
	//formdata
	data:=url.Values{}

	data.Add("firstname","Junaid")
	data.Add("lastName","Mumtaz")
	data.Add("emal","j@k.com")

	response,err:=	http.PostForm(myurl,data)
	handleNil(err)

	fmt.Println("response from server is",response)
	fmt.Println("response body from server is",response.Body)

	contentByte, err := io.ReadAll(response.Body)
	handleNil(err)

	fmt.Println(string(contentByte))

	defer response.Body.Close()

}

func handleNil(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

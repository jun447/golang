package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
)

func main(){
	// fmt.Println("webrequet")

	response,err := http.Get("https://example.com")
	handleNil(err)

	fmt.Println(response)
	contentByte,err:=io.ReadAll(response.Body)
	handleNil(err)
	fmt.Println(string(contentByte))
	defer response.Body.Close()


}

func handleNil(e error){
	if e!=nil{
		fmt.Println(e)
		panic(e)
	}
}
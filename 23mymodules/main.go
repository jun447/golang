package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Hello mod in golang")

	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000",r))
	
}

func greeter(){
	fmt.Println("Hey there mod Users")
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}

func handleNil(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test "}

var wg sync.WaitGroup
var mutex sync.Mutex
func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"http://localhost:8000",
		"http://localhost:8000/get",
		// "http://localhost:8000",
		"https://github.com",
		// "https://stackoverflow.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		// wg.Add(1)
		wg.Add(1)
	}
	wg.Wait() // wait for all goroutines to finish it will execute at the end
	fmt.Println(signals)

}
func getStatusCode(endpoint string) {


	res, err := http.Get(endpoint)
	handleNil(err)
	mutex.Lock()
	signals=append(signals, endpoint)
	mutex.Unlock()
	fmt.Printf("%d status code for URL %s \n", res.StatusCode, endpoint)
	defer wg.Done() // decrement the counter when the goroutine completes




	// for i := 0; i < 4; i++ {

	// 	res,err:=http.Get(endpoint)
	// 	handleNil(err)
	// 	fmt.Printf("%d status code for URL %s \n",res.StatusCode,endpoint)
	// 	defer wg.Done()  // decrement the counter when the goroutine completes
	// }
}

//	func greeter (s string){
//		for i := 0; i < 5 ; i++ {
//			time.Sleep(3* time.Second)
//			fmt.Println(s)
//		}
//	}
func handleNil(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

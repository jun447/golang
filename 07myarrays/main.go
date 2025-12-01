package main

import "fmt"
func main() {

	//fruitlist arrays
	var fruitList[4] string // [3]string{"apple","banana","grapes"}
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "orange"

	fmt.Println("fruitlist is ", fruitList)
	fmt.Println("lenght is ", len(fruitList))
}
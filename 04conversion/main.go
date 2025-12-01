package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	fmt.Println("Welcome to our saji house")
	fmt.Println("pleae rate our saji from 1 to 5")

	reader:= bufio.NewReader(os.Stdin)
	rating, _:= reader.ReadString('\n')
	fmt.Println("Your rating is ", rating)
	// type conversion
	// string to int
	// we have to trim the \n from string first
	var trimedRating string= rating[:len(rating)-1]
	// now convert
	var intRating int
	fmt.Sscanf(trimedRating, "%d", &intRating)
	fmt.Println("Converted rating is ", intRating)
	fmt.Printf("Type of intRating is %T \n", intRating)
	fmt.Printf("Type of rating is %T \n", rating)
	
		intRating+=1
		fmt.Println(intRating)

	//method 2
	// numRating,err:=strconv.ParseFloat(rating,64) error because of \n it also need trimming
	numRating,err:=strconv.ParseFloat(strings.TrimSpace(rating),64)

	if err!=nil{
		fmt.Println("Error occured", err)
	}else{
		fmt.Println("Converted rating using strconv is after adding 1 is", numRating+1)
		fmt.Printf("Type of numRating is %T \n", numRating)
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader:= bufio.NewReader(os.Stdin)
	welcome:="Nawa aya e sonya , nam btao please"
	fmt.Println(welcome)
	var name string
	// fmt.Scan(&name)
	// fmt.Println("Welcome", name)
	name, _= reader.ReadString('\n')
	fmt.Println("Welcome", name)



	// using buffreadermethod
	fmt.Println("PLease enter your uni name")

	// comma ok error ok
	// uniName, _:= reader.ReadString(',')
	uniName, err:= reader.ReadString('\n')
	fmt.Println("Your uni name is ", uniName)
	fmt.Printf("Type of uniName is %T \n", uniName)
    fmt.Println("Error is ", err)
	fmt.Printf("Error type is %T", err)


 }
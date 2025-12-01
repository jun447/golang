package main

import (
	"fmt"
	// "math"
)

const LoginToken string ="randomtoken123" // public variable capital letter make it public

func main()  {
	// fmt.Println("Hello, Variables!")
	// var myname string = "Junaid"
	// fmt.Printf("%s\n", myname)
	// fmt.Println(myname)
	// fmt.Printf("variable myname is of type %T \n", myname)

	// var smallVal uint8=255
	// fmt.Println(smallVal)
	// fmt.Printf("variable is of typ %T\n",smallVal)
    
	// fmt.Println("this is the max ize of float32",math.MaxFloat32)



	// var whatsinBeginning string
	// fmt.Println(whatsinBeginning)
	// fmt.Printf("variable whatsinBeginning is of type %T \n", whatsinBeginning)

	//implicit variable declaration
	var city = "Karachi"
	fmt.Println(city)
	fmt.Printf("variable city is of type %T \n", city)


	// no var stylr
	myCgpa := 3.75
	fmt.Println(myCgpa)
	fmt.Printf("variable myCgpa is of type %T \n", myCgpa)
}
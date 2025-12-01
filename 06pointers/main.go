package main

import "fmt"

func main() {
	// 	var age int = 25
	
	// // How do we find where 'age' lives?
	// fmt.Println(&age)  // & = "give me the address"
	// // Output: 0xc0000140a0  (hexadecimal address)
	// var ptr *int 
	// fmt.Println("value of this pointer is ", ptr) // nil

	myNumber :=23

	 var ptr = &myNumber // Go infers type *int

	// var ptr int= &myNumber //error
	fmt.Println("value of myNumber is ", myNumber)
	fmt.Println("address of myNumber is ", &myNumber)
	fmt.Println("value of pointer ptr is ", ptr)
	fmt.Println("value at the address stored in ptr is ", *ptr) // dereferencing
	*ptr=*ptr*78
	fmt.Println("new value of myNumber is ", myNumber)
}
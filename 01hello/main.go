package main

import (
	"fmt"
	"os"
)

func main(){
	// fmt.Println(os.Args)
	fmt.Println("Program arguments:")
	for i, arg := range os.Args {
		fmt.Printf("Arg %d: %s\n", i, arg)
	}
}
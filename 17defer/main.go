package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("worlds1")
	defer fmt.Println("worlds2")
	defer fmt.Println("worlds3")
	defer fmt.Println("worlds4")
	defer fmt.Println("worlds5")
	defer fmt.Println("worlds6")
	defer fmt.Println("worlds7")
	fmt.Println("Hello") //will print at first
	mydeferloop()
}

func mydeferloop() {
	for i := 0; i < 5; i++ {
		//  fmt.Print(i)
		defer fmt.Print(i)
	}
	fmt.Print("\n")
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dicNumber:=rand.Intn(6)
	fmt.Println(dicNumber)
	switch dicNumber {
	case 0:
		fmt.Println("zero is not allowed in dice")
	case 1:
		fmt.Println("You rolled a one")
	case 2:
		fmt.Println("You rolled a two")
	case 3:
		fmt.Println("You rolled a three")
	case 4:
		fmt.Println("You rolled a four")
	case 5:
		fmt.Println("You rolled a five")
	case 6:
		fmt.Println("You rolled a sixand can open your goti")
	default:
		fmt.Println("Unknown dice number")
	}
}
package main

import "fmt"

func main() {
	//creating slices
	// var fruitList=[]string{"Apple","tomato","banana"} // [3]string{"apple","banana","grapes"}
	// fmt.Println("fruitlist is ", fruitList)
	// fmt.Println("lenght is ", len(fruitList))
	// fmt.Printf("type of this slice is %T\n", fruitList)

	// fruitList=append(fruitList,"mango","kiwi")
	// fmt.Println("updated fruitlist is ", fruitList)
	// fmt.Println("updated lenght is ", len(fruitList))

	// //slicing slices
	// slicedFruit := fruitList[1:3]
	// fmt.Println("sliced fruitlist is ", slicedFruit)
	// //og list remains same
	// fmt.Println("original fruitlist is ", fruitList)

	// highScores := make([]int,4)

	// highScores[0]=23456
	// highScores[1]=56789
	// highScores[2]=34567
	// highScores[3]=45678

	// fmt.Println("highscores: ", highScores)
	// fmt.Println("length of it ", len(highScores))

	// highScores = append(highScores, 55555, 66666)
	// fmt.Println("updated highscores: ", highScores)
	// fmt.Println("length of highscores: ", len(highScores))
	// DELETING AN ELEMENT FROM A SLICE
	var courses = []string{"reactjs", "javascript", "html", "css", "nodejs", "expressjs"}
	fmt.Println("Original courses:", courses)
	// Output: [reactjs javascript html css nodejs expressjs]

	var index int = 2 // Want to delete element at index 2 (which is "html")

	// Step 1: courses[:index] creates a slice of first 2 elements
	fmt.Println("courses[:index]:", courses[:index])
	// Output: [reactjs javascript]

	// Step 2: courses[index+1:] creates a slice starting from index 3
	fmt.Println("courses[index+1:]:", courses[index+1:])
	// Output: [css nodejs expressjs]

	// Step 3: ... unpacks the second slice into individual elements
	// Without ...: append would try to append a SLICE to a SLICE (error!)
	// With ...: append gets individual elements (correct!)
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("After deletion:", courses)
	// Output: [reactjs javascript css nodejs expressjs]
	// Notice: "html" is removed!

}

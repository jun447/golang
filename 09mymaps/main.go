package main

import "fmt"

func main() {
	//creating maps
	// var myMap = make(map[string]int)
	languages := make(map[string]string)
	//how to declare it without := and use = this instead is it possible
	// var languages map[string]string
	// languages=make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"
	languages["go"] = "golang"
	languages["rs"] = "rust"
	languages["kt"] = "kotlin"

	fmt.Println(languages)

	fmt.Println("js stands for ", languages["js"])

	//delete an element from map
	delete(languages, "rb")
	fmt.Println("after deleting rb ", languages)

	fmt.Println(languages)

	//looping and dont u write it , i will write it now
	for key, value := range languages {

		fmt.Println(key, "stands for", value)
	}
	// var key string
	// var value string

	// for key,value=range languages{

	// 	fmt.Println(key, "stands for", value)
	// }

	// 	var value string

	// 	for value=range languages{

	//		fmt.Println("stands for", value)
	//	}
}

package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"os"

)

func main() {
	fmt.Println("The mighty file hanling")

	content := "dono jahan teri muhabbat main har ke\nwo ja raha hai koi shabe gham guzar ke"

	file, err := os.Create("./gazal.txt")
	if err != nil {
		fmt.Println("error occured", err)

	}
	// written, err := file.Write([]byte(content))  //method 1
	written, err := io.WriteString(file,content)  //method 2

	if err != nil {
		fmt.Println("error occured", err)
	}

	fmt.Println(written)
	defer file.Close()	
	readFile(file.Name())
}

func readFile(filename string){
	// contentByte,err:=ioutil.ReadFile(filename)
	contentByte,err:=os.ReadFile(filename)
		if err != nil {
		fmt.Println("error occured", err)
	}
	fmt.Println(string(contentByte))
}
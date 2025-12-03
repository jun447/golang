package main

import (
    "fmt"
    // "reflect"
)
func main(){
	greeter()
	// sum:=proAddres(1,2,3,4,5,6,7,8,9,10)
	var nums =[]int{1,2,3,4,5,6,7,8,9,10}
	// var nums =[10]int{1,2,3,4,5,6,7,8,9,10}
	// fmt.Printf("Type: %T\n", nums)  // []int (SLICE)
	// fmt.Println(reflect.TypeOf(nums).Kind())  // slice

	sum:=proAddres(&nums)
	fmt.Println("the sum is ",sum)
}

func greeter(){
	fmt.Println("hi fever stay away from me please")
}
// func proAddres(values... int) int {
// func proAddres(arr*[10] int) int {
func proAddres(arr*[] int) int {
	total:=0
	for _,v:=range(*arr){
		total=total+v
	}
	return total
}
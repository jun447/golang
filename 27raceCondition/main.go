package main

import (
	"fmt"
	"sync"
)
var  wg *sync.WaitGroup = &sync.WaitGroup{}
var mutex *sync.Mutex = &sync.Mutex{}
// var  wg sync.WaitGroup 

func main(){
	fmt.Println("Race Condition")

	var score=[]int{}

	wg.Add(3)
	
	go func(wg *sync.WaitGroup,mutex *sync.Mutex){
		fmt.Println("One R")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg,mutex)

	go func(wg *sync.WaitGroup,mutex *sync.Mutex){
		fmt.Println("Two R")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg,mutex)

	go func(wg *sync.WaitGroup,mutex *sync.Mutex){
		fmt.Println("Three R")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg,mutex)

	wg.Wait()
	fmt.Println("The Final value of Score is",score)
}
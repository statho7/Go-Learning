package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")
	
	fmt.Println("~*~")

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	
	fmt.Println("~*~")

	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}
package main

import (
	"fmt"
	"math"
)

func returnTrue() bool {
	return true
}

func main()  {
	// if false {
	// 	fmt.Println("The test is true")
	// } else {
	// 	fmt.Println("The test is yolo")
	// }

	// statePopulations := map[string]int{
	// 	"California":		39250017,
	// 	"Texas":			27862596,
	// 	"Florida":			20612439,
	// 	"New York":			19745286,
	// 	"Pennsylvania":		12802503,
	// 	"Illinois":			12801539,
	// 	"Ohio":				11614373,
	// }

	// if pop, ok := statePopulations["Florida"]; ok {
	// 	fmt.Println(pop)
	// }

	// number := 50
	// guess := 30
	// // if guess < 1 || returnTrue() || guess > 100 {
	// // 	fmt.Println("The guess must be between 1 and 100")
	// if guess < 1 {
	// 	fmt.Println("The guess must be greater than 1")
	// } else if guess > 100 {
	// 	fmt.Println("The guess must be less than 100")
	// } else {	
	// 	if guess < number {
	// 		fmt.Println("Too low!")
	// 	}
	// 	if guess > number {
	// 		fmt.Println("Too high!")
	// 	}
	// 	if guess == number {
	// 		fmt.Println("You got it!")
	// 	}
		
	// 	fmt.Println(number <= guess, number >= guess, number != guess)
	// }

	// fmt.Println(!true, !false)

	myNum := 0.123
	// if myNum == math.Pow(math.Sqrt(myNum), 2){
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1 ) < 0.001{
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	
	switch returnTrue() {
	case true:
		fmt.Println("Yolo")
		// fallthrough ---> bad, too bad be careful 
	case false:
		fmt.Println("Swag")		
	default:
		fmt.Println("Thug")
	}
	
	var i interface{}="1."
	switch i.(type) {
	case int:
		fmt.Println("i is int")
		// fallthrough ---> bad, too bad be careful 
	case float64:
		fmt.Println("i is float64")	
	case string:
		fmt.Println("i is string")		
		break
		fmt.Println("This won't print")
	default:
		fmt.Println("i is an helicopter")
	}
}
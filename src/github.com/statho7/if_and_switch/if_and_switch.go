package main

import (
	"fmt"
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

	number := 50
	guess := 30
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	}
	if guess >= 1 && guess <= 100 {	
		if guess < number {
			fmt.Println("Too low!")
		}
		if guess > number {
			fmt.Println("Too high!")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	fmt.Println(!true, !false)

	switch returnTrue() {
	case true:
		fmt.Println("Yolo")
	case false:
		fmt.Println("Swag")
		
	}
}
package main

import (
	"fmt"
)

func main() {
	// count := 10
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	// for i, j := 0, 0; i < count; i, j = i + 1, j + 2 {
	// 	fmt.Println(i, j)
	// }

	// for i := 0; i < count; i++ {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// i := 0
	// for ; i < count; i++ {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// i := 0
	// for ; i < count;  {
	// 	i++
	// 	if i % 2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// //infinite loop
	// i := 0
	// for i < count {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// i := 0
	// for i < count {
	// 	if i % 9 == 0 {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// }

	// for i := 0; i < count; i++ {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i)
	// 		continue
	// 	}
	// }

	// Loop:
	// for i := 0; i < count; i++ {
	// 	for j := 0; j < count; j++ {
	// 		if i % 2 == 0 {
	// 			fmt.Println(i * j)
	// 		}
	// 		if (i * j) > 30{
	// 			break Loop
	// 		}
	// 	}
	// }

	// s := []int{1, 2, 3}
	// for k, v := range s {
	// 	fmt.Println(k, v)
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

	// for k, v := range statePopulations {
	// 	fmt.Println(k, v)
	// }

	s := "Hello Go"
	for k, v := range s {
		fmt.Println(k, v)
	}
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}

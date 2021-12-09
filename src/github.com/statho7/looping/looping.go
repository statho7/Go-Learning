package main

import (
	"fmt"
)

func main()  {
	count := 10
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

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i % 2 == 0 {
				fmt.Println(i * j)
			}			
		}
	}
}
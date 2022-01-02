package main

import (
	"fmt"
)

// func main()  {
// 	sayMessage("Hello Go!")
// }

// func sayMessage(msg string)  {
// 	fmt.Println(msg)
// }

// func main()  {
// 	for i := 0; i < 5; i++ {
// 		sayMessage("Hello Go!", i)
// 	}
// }

// func sayMessage(msg string, idx int)  {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is", idx)
// }

// func main()  {
// 	greeting := "Hello"
// 	name := "Stacey"
// 	sayGreeting(greeting, name)
// }

// func sayGreeting(greeting, name string)  {
// 	fmt.Println(greeting, name)
// 	name = "Ted"
// 	fmt.Println(name)
// }

func main()  {
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
}

func sayGreeting(greeting, name *string)  {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}
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

// func main()  {
// 	greeting := "Hello"
// 	name := "Stacey"
// 	sayGreeting(&greeting, &name)
// 	fmt.Println(name)
// }

// func sayGreeting(greeting, name *string)  {
// 	fmt.Println(*greeting, *name)
// 	*name = "Ted"
// 	fmt.Println(*name)
// }

// func main()  {
// 	sum("The sum is ", 1,2,3,4,5)
// }

// func sum(msg string, values ...int)  {
// 	fmt.Println(values)
// 	result := 0
// 	for _,v := range values {
// 		result += v
// 	}
// 	fmt.Println(msg, result)
// }

// func main()  {
// 	s := sum(1,2,3,4,5)
// 	fmt.Println("The sum is ", s)
// }

// func sum(values ...int) int {
// 	fmt.Println(values)
// 	result := 0
// 	for _,v := range values {
// 		result += v
// 	}

// 	return result
// }

// func main()  {
// 	s := sum(1,2,3,4,5)
// 	fmt.Println("The sum is ", *s)
// }

// func sum(values ...int) *int {
// 	fmt.Println(values)
// 	result := 0
// 	for _,v := range values {
// 		result += v
// 	}

// 	return &result
// }

// func main()  {
// 	s := sum(1,2,3,4,5)
// 	fmt.Println("The sum is ", s)
// }

// func sum(values ...int) (result int) {
// 	fmt.Println(values)
// 	for _,v := range values {
// 		result += v
// 	}

// 	return
// }

// func main()  {
// 	d,err := divide( 5.0, 0.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Can not divide by zero.")
// 	}
// 	return a / b, nil
// }

func main()  {
	// func ()  {
	// 	msg := "Hello Go!"
	// 	fmt.Println(msg)
	// }()

	// for i := 0; i < 5; i++ {
	// 	func (i int)  {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	
	// var f func() = func() {
	// 	msg := "Hello Go!"
	// 	fmt.Println(msg)
	// }

	// f := func ()  {
	// 	msg := "Hello Go!"
	// 	fmt.Println(msg)
	// }
	// f()

	// var divide func( float64, float64) (float64, error)
	// divide = func(a, b float64) (float64, error) {
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("Can not divide by zero.")
	// 	}
	// 	return a / b, nil
	// }
	// d,err := divide( 5.0, 0.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(d)

	g := greeter{
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
	fmt.Println("The new name is: ", g.name)
}

type greeter struct {
	greeting string
	name string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

// func (g *greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = ""
// }
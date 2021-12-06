package main

import (
    "fmt"
)

// func main() {
// 	// n := 1 == 1
// 	// m := 1 == 2
// 	// fmt.Printf("%v, %T\n",n,n)
// 	// fmt.Printf("%v, %T\n",m,m)

// 	// var n bool
// 	// fmt.Printf("%v, %T\n",n,n)

// 	// var n int
// 	// n := 42
// 	var n uint16 = 42
// 	fmt.Printf("%v, %T\n",n,n)
// }

// func main() {
// 	a := 10
// 	b := 3
// 	fmt.Println(a + b)
// 	fmt.Println(a - b)
// 	fmt.Println(a * b)
// 	fmt.Println(a / b)
// 	fmt.Println(a % b)
// }

// func main() {
// 	var a int = 10
// 	var b int8 = 3
// 	// fmt.Println(a + b) Error
// 	fmt.Println(a + int(b))
// 	// fmt.Println(a - b)
// 	// fmt.Println(a * b)
// 	// fmt.Println(a / b)
// 	// fmt.Println(a % b)
// }


// // bit operations
// func main() {
// 	a := 10
// 	b := 3
// 	fmt.Println(a & b)
// 	fmt.Println(a | b)
// 	fmt.Println(a ^ b)
// 	fmt.Println(a &^ b)
// }

// func main()  {
// 	a := 7
// 	fmt.Println(a << 3)
// 	fmt.Println(a >> 3)
// }

//floating point numbers
func main()  {
	n := 3.14
	fmt.Printf("%v, %T\n",n,n)
	n = 13.7e72
	fmt.Printf("%v, %T\n",n,n)
	n = 2.1E14
	fmt.Printf("%v, %T\n",n,n)
}
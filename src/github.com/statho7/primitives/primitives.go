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


// bit operations
func main() {
	a := 10
	b := 3
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
}
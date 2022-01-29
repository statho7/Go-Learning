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

// //floating point numbers
// func main()  {
// 	n := 3.14
// 	fmt.Printf("%v, %T\n",n,n)
// 	n = 13.7e72
// 	fmt.Printf("%v, %T\n",n,n)
// 	n = 2.1E14
// 	fmt.Printf("%v, %T\n",n,n)
// }

// func main()  {
// 	a := 10.2
// 	b := 3.7
// 	fmt.Println(a + b)
// 	fmt.Println(a - b)
// 	fmt.Println(a * b)
// 	fmt.Println(a / b)
// }

// imaginary
// func main()  {
// 	// var n complex64 = 2i
// 	// var n complex64 = 1 + 2i
// 	// fmt.Printf("%v, %T\n",n,n)
// 	// a := 1 + 2i
// 	// b := 2 + 5.2i
// 	// fmt.Println(a + b)
// 	// fmt.Println(a - b)
// 	// fmt.Println(a * b)
// 	// fmt.Println(a / b)

// 	// var n complex64 = 1 + 2i
// 	// var n complex128 = 1 + 2i
// 	// fmt.Printf("%v, %T\n",real(n),real(n))
// 	// fmt.Printf("%v, %T\n",imag(n),imag(n))

// 	var n complex128 = complex(5,12)
// 	fmt.Printf("%v, %T\n",n,n)
// }

// func main() {
// 	// s:= "this is string"
// 	// fmt.Printf("%v, %T\n",s,s)

// 	// s:= "this is string"
// 	// fmt.Printf("%v, %T\n",s[2],s[2])

// 	// s:= "this is string"
// 	// s2 := "this is also a string"
// 	// fmt.Printf("%v, %T\n",s + s2,s + s2)

// 	s:= "this is string"
// 	b := []byte(s)
// 	fmt.Printf("%v, %T\n",b,b)
// }

// runes
func main() {
	// r:= 'a'
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}

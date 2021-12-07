package main

import (
	"fmt"
	// "math"
)

// func main()  {
// 	// const myConst int = 42
// 	// const myConst float64 = math.Sin(1.57) // Error
// 	// fmt.Printf("%v, %T\n", myConst, myConst)
// }

// func main()  {
// 	const a int = 14
// 	const b string = "foo"
// 	const c float32 = 3.14
// 	const d bool = true
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%v\n", d)
// }

// const a int16 = 27

// func main()  {
// 	// const a int = 14
// 	fmt.Printf("%v, %T\n", a, a)
// }

// func main()  {
// 	const a int = 42
// 	const b int = 27
// 	fmt.Printf("%v, %T\n", a + b, a + b)
// }

// func main()  {
// 	const a int = 42
// 	const b int16 = 27
// 	fmt.Printf("%v, %T\n", a + b, a + b) //Error
// }

// func main()  {
// 	const a = 42
// 	const b int16 = 27
// 	fmt.Printf("%v, %T\n", a + b, a + b)
// }

// enumerated constant

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )
const (
	a = iota
	b
	c
)

const (
	a2 = iota
)

func main()  {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", a2, a2)
}

// enumerated constant example app
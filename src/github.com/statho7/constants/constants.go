package main

import (
	"fmt"
	// "math"
)

func main()  {
	// const myConst int = 42
	// const myConst float64 = math.Sin(1.57) // Error
	// fmt.Printf("%v, %T\n", myConst, myConst)

	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}
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
// const (
// 	a = iota
// 	b
// 	c
// )

// const (
// 	a2 = iota
// )

// func main()  {
// 	fmt.Printf("%v, %T\n", a, a)
// 	fmt.Printf("%v, %T\n", b, b)
// 	fmt.Printf("%v, %T\n", c, c)
// 	fmt.Printf("%v, %T\n", a2, a2)
// }

// enumerated constant examples

// const (
// 	// errorSpecialist = iota
// 	// _ = iota
// 	_ = iota + 5
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )

// func main()  {
// 	// var specialistType int = catSpecialist
// 	// fmt.Printf("%v\n", specialistType == catSpecialist)
// 	// var specialistType int = dogSpecialist
// 	// fmt.Printf("%v\n", specialistType == dogSpecialist)
// 	// var specialistType int
// 	fmt.Printf("%v\n", dogSpecialist)
// }

// const (
// 	_ = iota // ignore first value by assigning to blank identifier
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

// func main()  {
// 	fileSize := 4000000000.
// 	fmt.Printf("%.2fGB", fileSize/GB)
// }

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main()  {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin ? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is HQ ? %v", isHeadquarters & roles == isHeadquarters)
}
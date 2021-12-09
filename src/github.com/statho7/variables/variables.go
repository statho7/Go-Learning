package main

// import "fmt"

// var j int = 27

// func main() {
//     // var i int
//     // i=42
//     // var j int = 27
//     fmt.Printf("%v, %T", j, j)
//     // k := 99
//     // fmt.Printf("%v, %T", k, k)
// }
//
// var (
//     var actorName string = "Elizabeth Sladen"
//     var companion string = "Sarah Jane Smith"
//     doctorNumber int = 3
//     season int = 11
// )

// var (
//     counter int = 0
// )

// func main() {
    
// }

// var j int = 27

// func main() {
//     fmt.Printf("%v, %T\n", j, j)
//     var j int = 44
//     fmt.Printf("%v, %T", j, j)
// }

// var J int = 27

// func main() {
//     var i int = 42
//     fmt.Printf("%v, %T\n", i, i)
    
//     var j float32
//     j = float32(i)
//     fmt.Printf("%v, %T\n", j, j)
// }


import (
    "fmt"
    "strconv"
)

func main() {
    var i int = 42
    fmt.Printf("%v, %T\n", i, i)
    
    var j string = strconv.Itoa(i)
    fmt.Printf("%v, %T\n", j, j)
}
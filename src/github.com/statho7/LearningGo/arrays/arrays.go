package main

import "fmt"

func main(){
	var x1 = [...]int{10, 20, 30}
	fmt.Println(x1)

	var x2 = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x2 == y)

	var x3 [2][3]int
	fmt.Println(x3)
}
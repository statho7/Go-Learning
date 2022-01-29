package main

import "fmt"

func main(){
	// If you have a sparse array (an array where most elements are set to their zero value), you can specify only the indices with values in the array literal:
	var x = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x)
}
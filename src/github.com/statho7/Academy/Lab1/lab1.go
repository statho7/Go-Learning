package main

import (
	"fmt"
	"strconv"
)

func product(array [5]int) int {

	product := array[0]
	i := 1
	length := (len(array) - 1)

	for i <= length {
		product = product * array[i]
		i = i + 1
	}

	return product
}

func slicemania(mySlice []string) (int, string) {

	length := (len(mySlice) - 1)
	var total int
	total = 1
	myString := ""
	for i := 0; i <= length; i++ {
		myInt, b := strconv.Atoi(mySlice[i])
		if b != nil {
			myString += mySlice[i]
		} else if myInt%2 == 0 {
			total = total * myInt
		}
		// fmt.Println(myInt,b)
	}

	return total, myString
}

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := []string{"1", "4", "293", "4", "9", "hello", "sunshine"}

	fmt.Println(myArray)
	fmt.Println(mySlice)

	res1 := product(myArray)
	fmt.Println(res1)

	res2, res3 := slicemania(mySlice)
	fmt.Println(res2, res3)
}

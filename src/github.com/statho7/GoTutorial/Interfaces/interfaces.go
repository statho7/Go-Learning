package main

import (
	"fmt"
)

// func main() {
// 	var w Writer = ConsoleWriter{}
// 	w.Write([]byte("Hello Go!"))
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct {}

// func(cw ConsoleWriter) Write(data []byte) (int, error){
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

func main(){
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface{
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int{
	*ic++
	return int(*ic)
}
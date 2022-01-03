package main

import (
	"fmt"
	"sync"
)

// var wg = sync.WaitGroup{}

// func main() {
// 	// go sayHello()
// 	var msg = "Hello"
// 	wg.Add(1)
// 	go func(msg string){
// 		fmt.Println(msg)
// 		wg.Done()
// 	}(msg)
// 	msg = "Goodbye"
// 	wg.Wait()
// }

// // func sayHello() {
// // 	fmt.Println("Hello")
// // }

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()		
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello #v%\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
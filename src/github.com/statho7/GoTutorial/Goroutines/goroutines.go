package main

import (
	"fmt"
	"runtime"
	// "sync"
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

// var wg = sync.WaitGroup{}
// var counter = 0
// var m = sync.RWMutex{}

// func main() {
// 	runtime.GOMAXPROCS(1000)
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(2)
// 		m.RLock()
// 		go sayHello()
// 		m.Lock()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Println("Hello #%v\n", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	m.Unlock()
// 	wg.Done()
// }

func main() {
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v", runtime.GOMAXPROCS(-1))
}

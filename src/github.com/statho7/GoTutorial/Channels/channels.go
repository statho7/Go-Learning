package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func(){
// 		i := <- ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
// 	go func(){
// 		i := 42
// 		ch <- i
// 		i = 27
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

// func main() {
// 	ch := make(chan int)
// 	for j := 0; j < 5; j++ {
// 		wg.Add(2)
// 		go func(){
// 			i := <- ch
// 			fmt.Println(i)
// 			wg.Done()
// 		}()
// 		go func(){
// 			i := 42
// 			ch <- i
// 			i = 27
// 			wg.Done()
// 		}()
// 		wg.Wait()
// 	}
// }

// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <- chan int){
// 		for i := range ch{
// 			fmt.Println(i)
// 		}
// 		// i = <- ch
// 		// fmt.Println(i)
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int){
// 		ch <- 42
// 		ch <- 27
// 		close(ch)
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <- chan int){
// 		for {
// 			if i, ok := <- ch; ok {
// 				fmt.Println(i)
// 			} else {
// 				break
// 			}
// 		}
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int){
// 		ch <- 42
// 		ch <- 27
// 		close(ch)
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// import (
// 	"fmt"
// 	"time"
// )

// const (
// 	logInfo = "INFO"
// 	logWarning = "WARNING"
// 	logError = "ERROR"
// )

// type logEntry struct {
// 	time time.Time
// 	severity string
// 	message string
// }

// var logCh = make(chan logEntry, 50)

// func main() {
// 	go logger()
// 	defer func(){
// 		close(logCh)
// 	}()
// 	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

// 	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
// 	time.Sleep(100 * time.Millisecond)
// }

// func logger() {
// 	for entry := range logCh {
// 		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
// 	}
// }

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main() {
	go logger()
	defer func() {
		close(logCh)
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
			// default:
			// 	fmt.Println("Default")
		}
	}
}

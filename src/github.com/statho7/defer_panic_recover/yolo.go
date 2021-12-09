package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("start")
// 	fmt.Println("middle")
// 	fmt.Println("end")
	
// 	fmt.Println("~*~")

// 	fmt.Println("start")
// 	defer fmt.Println("middle")
// 	fmt.Println("end")
	
// 	fmt.Println("~*~")

// 	defer fmt.Println("start")
// 	defer fmt.Println("middle")
// 	defer fmt.Println("end")
// }

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
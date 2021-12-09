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

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	robots, err := ioutil.ReadAll(res.Body)
// 	defer res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%s", robots)
// }

// import (
// 	"fmt"
// )

// func main() {
// 	// a := "start"
// 	// defer fmt.Println(a)
// 	// a = "end"
	
// 	// a, b := 1, 0
// 	// ans := a / b
// 	// fmt.Println(ans)
	
// 	fmt.Println("start")
// 	panic("Something bad happened")
// 	fmt.Println("end")
// }

// import (
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello Go!"))
// 	})
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// import (
// 	"fmt"
// 	"log"
// )

// func main() {	
// 	// fmt.Println("start")
// 	// defer fmt.Println("this was deferred")
// 	// panic("Something bad happened")
// 	// fmt.Println("end")
	
// 	fmt.Println("start")
// 	defer func() {
// 		if err:= recover(); err != nil {
// 			log.Println("Error:", err)
// 		}
// 	}()
// 	panic("Something bad happened")
// 	fmt.Println("end")
// }

import (
	"fmt"
	"log"
)

func main() {	
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic!")
	defer func() {
		if err:= recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Something bad happened")
}
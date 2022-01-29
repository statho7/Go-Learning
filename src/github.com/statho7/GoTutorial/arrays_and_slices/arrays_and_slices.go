package main

import (
	"fmt"
)

func main() {
	// grade1 := 97
	// grade2 := 85
	// grade3 := 93
	// fmt.Printf("Grades: %v, %v, %v", grade1, grade2, grade3)

	// grades := [3]int{97,85,93}
	// fmt.Printf("Grades: %v", grades)

	// grades := [...]int{97,85,93}
	// fmt.Printf("Grades: %v", grades)

	// var students [3]string
	// fmt.Printf("Students: %v\n", students)
	// students[0] = "Lisa"
	// students[2] = "Ahmed"
	// students[1] = "Arnold"
	// fmt.Printf("Students: %v\n", students)
	// fmt.Printf("Students #1: %v\n", students[1])
	// fmt.Printf("Number of students: %v", len(students))

	// var identityMatrix [3][3]int = [3][3]int { [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1}}
	// fmt.Println(identityMatrix)

	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1,0,0}
	// identityMatrix[0] = [3]int{0,1,0}
	// identityMatrix[0] = [3]int{0,0,1}
	// fmt.Println(identityMatrix)

	// a := [...]int{1, 2, 3}
	// b := a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	// a := [...]int{1, 2, 3}
	// b := &a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	// a := []int{1, 2, 3}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{1, 2, 3}
	// b := a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// fmt.Println(b)
	// fmt.Printf("Length: %v\n", len(b))
	// fmt.Printf("Capacity: %v\n", cap(b))

	// // a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]
	// // a[5] = 42
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// // a := make([]int, 3, 100)
	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 2, 3, 4, 5)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// // a = append(a, []int{2, 3, 4, 5})
	// // fmt.Println(a)
	// // fmt.Printf("Length: %v\n", len(a))
	// // fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, []int{2, 3, 4, 5}...)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{1, 2, 3, 4, 5}
	// // b := a[1:]
	// // b := a[:len(a)-1]
	// fmt.Println(a)
	// b := append(a[:2], a[3:]...)
	// fmt.Println(b)
	// fmt.Println(a)a[4] = 100

	var a [5]int
	fmt.Println("emp:", a)
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

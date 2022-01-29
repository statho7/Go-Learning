package main

import (
	"fmt"
	"math"
)

type Circle struct {
	diameter, radius int
}

func (r *Circle) area() float64 {
	pi := math.Pi
	return pi * math.Pow(float64(r.radius), 2)
}

func (r Circle) circumference() float64 {
	pi := math.Pi
	return float64(2) * pi * float64(r.radius)
}

func main() {
	r := Circle{diameter: 10, radius: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.circumference())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.circumference())
}

package main

import (
	"fmt"
	"math"
)

type cirle struct {
	radius float64
}
type rectangle struct {
	width, heigh float64
}

type shape interface {
	area() float64
	perimeter() float64
}

func (r cirle) area() float64 {
	return 2 * math.Pi * math.Pow(r.radius, 2)
}
func (r cirle) perimeter() float64 {
	return 2 * math.Pi * r.radius
}

func (r rectangle) area() float64 {
	return r.heigh * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.heigh + r.width)
}
func (v cirle) Volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(v.radius, 3)
}

func print(s shape) {
	fmt.Printf("shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("perimeter: %v\n", s.perimeter())
}

func main() {
	var s shape
	fmt.Printf("%T\n", s)

	fmt.Println("--------------------------------")

	ball := cirle{radius: 5.}
	s = ball
	print(s)
	fmt.Printf("Type of  s %T\n", s)

	fmt.Println("--------------------------------")

	room := rectangle{heigh: 5., width: 4.}
	s = room
	print(s)
	fmt.Printf("Type of s %T\n", s)

	fmt.Println("--------------------------------")

	s = cirle{radius: 2.5}
	fmt.Printf("Type of  s %T\n", s)
	val, ok := s.(cirle)
	if ok == true {
		fmt.Printf("val volume is %v\n", val.Volume())
	}

}

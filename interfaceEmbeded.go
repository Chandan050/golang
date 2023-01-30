package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type object interface {
	volume() float64
}

//geometry is embedded shape and object interface
type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge  float64
	Color string
}

func (c cube) area() float64 {
	return 6 * math.Pow(c.edge, 2)
}
func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}
func (c cube) getColor() string {
	return c.Color
}
func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("area %v, volume %v\n", a, v)

}

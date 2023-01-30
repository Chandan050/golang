package main

import "fmt"

type minute = uint

func main() {
	type speed int

	var x int = 56
	var y speed = 78
	fmt.Println(speed(x) - (y))
	fmt.Printf("%v %v \n", x, y)

	var a uint8
	var b byte
	fmt.Println(a == b)

	var t1 minute = 10
	var u uint = t1
	fmt.Printf("%T%T", t1, u)

	if t1 <= 10 {
		fmt.Println("njh")
	}

}

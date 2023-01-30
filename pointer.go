package main

import "fmt"

func main() {

	name := "hey"
	ptr := &name
	fmt.Println(*ptr, ptr)

	var ptr1 *float64
	_ = ptr1
	p := new(int)
	x := 100
	p = &x
	*p = 90

	fmt.Printf("p is of type %T wuth a value of %v %v %d\n", p, p, *p, x)

	//* comparing pointers *//
	var p2 *int
	fmt.Printf(" pointer to int/t%#v\n", p2)
	fmt.Println(*p == x)

}

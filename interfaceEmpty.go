package main

import "fmt"

type emptyInterface interface{}

type person struct {
	Info interface{}
}

func main() {

	var empty interface{}
	empty = 5
	fmt.Println(empty)
	fmt.Printf("%T\n", empty)

	empty = "go"
	fmt.Println(empty)
	fmt.Printf("%T\n", empty)

	empty = []int{1, 2, 3, 4}
	fmt.Println(empty)
	fmt.Printf("%T\n", empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.Info = "your name"
	you.Info = 80
	you.Info = []float64{4.5, 5., 7.}
	fmt.Println(you.Info)

}

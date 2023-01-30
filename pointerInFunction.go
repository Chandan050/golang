package main

import (
	"fmt"
	f "fmt"
)

func change(a *int) *float64 {
	*a = 100
	b := 5.5
	return &b

}
func changeValue(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.76
	name = "mobile phone"
	sold = true

}
func changeValueByPointers(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.76
	*name = "mobile phone"
	*sold = true
	fmt.Println("after the func call", *quantity, *price, *name, *sold)

}

type values struct {
	quantity int
	price    float64
	name     string
	sold     bool
}

func changeValues(v *values) {
	v.quantity = 900
	v.price = 90.879
	v.name = "helicaptor"
	v.sold = false

}

func main() {

	// pointers in go
	x := 8
	p := &x

	f.Printf("the value of x is %#v and addres of x/p is %v\n", x, p)
	a := change(p)

	f.Printf("the value of x is %#v and addres of a is %v\n", a, *a)

	q, pr, n, s := 7, 900.6, "plane", true
	fmt.Println("befoe func call", q, pr, n, s)
	changeValueByPointers(&q, &pr, &n, &s)

	gift := values{
		price:    90171313,
		quantity: 2,
		name:     "metro",
		sold:     true,
	}
	fmt.Println("befor sending to function")

	// to changeValue function
	changeValue(gift.quantity, gift.price, gift.name, gift.sold)

	// values after the func return
	f.Println(gift.quantity, gift.price, gift.name, gift.sold)

	// passing to change that values by pointers
	fmt.Println("------------------------------------")

	fmt.Println(gift)
	changeValueByPointers(&gift.quantity, &gift.price, &gift.name, &gift.sold)
	f.Println(gift)

	f.Println("--------------------------------------")
	// POITER TO THE SLICE

	prices := []int{1, 2, 3}
	fmt.Println(prices)

	pri := chnageSlice(&prices)
	fmt.Println(pri)

	// POINTER TO MAP
	f.Println("-----------------------------------------")
	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println("myMap after calling changeMap:", myMap)

	f.Println("-----------------------------------------")
	x1 := 1
	x2 := f1(x1)
	f.Println(x2)
}

func chnageSlice(v *[]int) []int {
	b := []int{}
	for a := range *v {
		b = append(b, a+1)
	}
	return b

}

// MAP POINTER
func changeMap(m map[string]int) {
	m["a"] = 30
	m["b"] = 20
	m["c"] = 10

}
func f1(x int) int {
	return x + 1
}

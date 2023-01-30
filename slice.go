package main

import (
	"fmt"
	f "fmt"
	"unsafe"
)

func main() {
	cities := []string{"hey", "hello", "nay"}
	fmt.Println(cities)
	f.Printf("%T %v \n", cities, len(cities))

	s := []int{1, 2, 3, 4}
	k := [4]int{1, 2, 3, 4}
	f.Println(unsafe.Sizeof(s))
	f.Println(unsafe.Sizeof(k))

	d := []int{5, 6, 7}
	s = append(s, d...)
	f.Println("sorce", s, "destiny", d)
	f.Println(s, d, copy(s, d))

	bar := []string{}
	rab := cities[0:2]
	bar = append(bar, cities[0:2]...)
	f.Println(cities, bar)
	bar[0] = "bye"
	f.Println(cities, bar, rab)
	cities[0] = "baby"
	f.Println(cities, bar, rab)
	f.Printf("%p %v %p", &cities[0], bar, &rab[0])

}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int
	fmt.Println("enter th")
	fmt.Scanln("&input", input)
	s := []int{4, 5, 2, 1, 3}
	sort.Ints(s)
	fmt.Println(s)
}

ō

package main

import (
	"fmt"
	f "fmt"
)

func main() {
	a := map[string]string{"A": "1"}
	b := map[string]string{"a": "1"}

	s1 := f.Sprintf("%s", a)
	s2 := f.Sprintf("%s", b)
	f.Println(s1 == s2)
	f.Printf("%T\t%T\t\n", s1, s2)

	c := a
	a["b"] = "3"
	fmt.Println(a, c)

	delete(a, "A")
	fmt.Println(a, c)

}

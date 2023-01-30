package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func myFunc(s string) int {

	var str string
	var sum int
	for i := 1; i <= 3; i++ {
		str += s
		n, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	return sum

}
func searchItem(s []string, v string) (string, bool) {
	str := fmt.Sprintf("%s", s)
	return "the search items results: ", strings.Contains(str, v)

}

func main() {
	a := myFunc("5")
	fmt.Println(a)
	animals := []string{"lion", "tiger", "bear"}
	result, v := searchItem(animals, "bear")
	fmt.Println(result, v) // => true
	result, v = searchItem(animals, "pig")
	fmt.Println(result, v)
}

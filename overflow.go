package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 8                  // int value
	y := 9.1                // float value
	fmt.Println(x * int(y)) // * of int & float

	var xy float32 = 67.98
	var yx float64 = 67.98
	fmt.Println((xy*float32(yx))/xy, xy*float32(yx))

	s := string(99)
	fmt.Println(s)

	var mystr = fmt.Sprintf("%f", 99.87)
	var mystr1 = fmt.Sprintf("%d", 99)

	fmt.Printf("%T %v \n", mystr, mystr)
	fmt.Printf("%T %v", mystr1, mystr1)

	kk := "3.1422"
	var f1, _ = strconv.ParseFloat(kk, 64)
	fmt.Println("--------")
	fmt.Println(f1)

	fmt.Println("------------------------")
	i, err := strconv.Atoi("-89")
	if err != nil {
		fmt.Println(err)
	}
	s2 := strconv.Itoa(78)
	fmt.Println(i, s2)
	fmt.Printf("type of i %T and type of s2 %T\n", i, s2)

}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("os.Args", os.Args[0])
	fmt.Println("os.Args", os.Args[1])
	fmt.Println("os.Args", os.Args)
	fmt.Println("os.Args", len(os.Args))
	fmt.Println("----------------------")
	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("the result %T %v ", result, result)
	_ = err
}

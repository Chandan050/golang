package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	pep := [5]int{1, 2, 3, 4, 5}
	pop := [2]int{2, 6}

loop:
	for index, value := range pep {
		for _, val := range pop {
			if val == value {
				fmt.Println("the commomn friends are", index, val)
				break loop
			}
		}
	}

term:
	if i <= 10 {
		switch {
		case i%2 == 0:
			fmt.Println("even number", i)
		case i%2 != 0:
			fmt.Println("odd number", i)
		default:
			fmt.Println("etered wrong number")
		}
		i++
		goto term

	}
	hour := time.Now().Hour()

	fmt.Println("number of hours", hour, time.Now().Minute())

}

package main

import (
	"fmt"
)

func generate(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 2; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	// close(out)
	return out
}
func prime(ch chan int) chan int {
	in := make(chan int)
	go func() {
		for n := range ch {
			c := 0
			for i := 2; i <= (n / 2); i++ {
				if n%i == 0 {
					c += 1
				}
			}
			if c == 0 {
				in <- n
			}
		}
		close(in)
	}()
	return in
}

func main() {
	n := 13
	value := generate(n)
	fin := prime(value)

	for n := range fin {
		fmt.Printf("%d ", n)
	}
}

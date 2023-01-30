package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	n := 4
	go func() {
		for i := 1; i <= n; i++ {
			c <- i
		}
		close(c)
	}()
	fmt.Println("", c)
	var sum int = 1
	for n := range c {
		sum *= n
	}
	fmt.Printf("foctorial of: %v is :%d \n", n, sum)
	fmt.Println("-----------------------------")

	k := gen()
	e := foctorial(k)
	for i := range e {
		fmt.Println(i)
	}

	fmt.Println("---------------------------------")
	m := fanin(boring("joe"), boring("anna"))
	for i := 1; i <= 10; i++ {
		fmt.Println(<-m)
	}

}
func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
func foctorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out

}
func fact(n int) int {
	var sum int = 1
	for i := n; i >= 1; i-- {
		sum *= i
	}
	return sum
}
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 10; i++ {
			c <- fmt.Sprintf("%s  %v ", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

//FAN IN
func fanin(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

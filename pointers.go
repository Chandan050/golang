package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	// var wg sync.WaitGroup

	// wg.Add(2)
	for i := 0; i <= 10; i++ {
		// wg.Wait()
		go func() {
			for i := 0; i <= 10; i++ {
				c <- i
			}
			done <- true
		}()
		// wg.Done()

	}
	go func() {
		for i := 0; i <= 10; i++ {
			<-done
		}
		close(c)
		// wg.Done()
	}()
	fmt.Println(c)
	a := 0
	for n := range c {
		fmt.Println(n, "   ", a)
		a++
	}

	fmt.Println("--------------------------------------")
	c1 := increament("meow")
	c2 := increament("bow")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("final counter", <-c3+<-c4)

}

func increament(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}
func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

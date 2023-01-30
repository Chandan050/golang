package main

import (
	"fmt"
	"math"
	"strconv"
)

func armatstrong(num int) int {
	pow := len(strconv.Itoa(num))
	temp := num
	result := 0
	for temp > 0 {
		rem := temp % 10
		result += int(math.Pow(float64(rem), float64(pow)))
		temp = temp / 10

	}
	return result
}

func palindrome(num int) int {
	temp := num
	result := 0
	for temp > 0 {
		rem := temp % 10
		result = result*10 + rem
		temp /= 10

	}
	return result
}

func prime(num int) {
	val := 0
	for i := 2; i <= (num/2)+1; i++ {
		if num%i == 0 {
			val += 1
			break
		}
	}
	if val == 0 || num == 1 || num == 2 {
		fmt.Printf("%d number is prime", num)
	} else {
		fmt.Printf("%d number is not prime", num)
	}
}

func fibonacci(num int) {
	a := 0
	b := 1
	fmt.Printf("%d %d ", a, b)
	for num-2 > 0 {
		var c = a + b
		a = b
		b = c
		fmt.Printf("%d ", c)
		num -= 1
	}
}

func main() {
	var sa int
	sa, _ = fmt.Scanf()
	va := armatstrong(153)
	if va == sa {
		fmt.Println(sa, " is amdtrong number")
	} else {
		fmt.Println(sa, " is not amstrong number")
	}

	fmt.Println(va)
	pa := palindrome(1221)
	fmt.Println(pa)
	prime(6)
	fmt.Println()
	fmt.Println("fibonacci series :")
	fibonacci(5)
}

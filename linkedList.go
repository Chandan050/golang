package main

import (
	"fmt"
	"os"
	"strconv"
)

// struct type creating for values to assign in future
type linkedlist struct {
	number int
	next   *linkedlist
}

// to insrt value into the linked list by using struct type
func (node *linkedlist) insert(num int) {
	var temp = &linkedlist{}
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp
		lastNode := temp
	} else {
		var p = &linkedlist{}
		p = node

		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}
}
func (node *linkedlist) insertLast()

// to display value in linked list -to take input to function(value refence)
func (node *linkedlist) display() {
	var p = &linkedlist{}
	p = node.next

	for p != nil {
		fmt.Printf("%d-->", p.number)
		p = p.next
	}
	fmt.Println()

}

func (node *linkedlist) deletefirst() *linkedlist {
	var p = node
	node = p.next
	return p
}
func (node *linkedlist) deletelast() {
	var p *linkedlist = node
	for p.next != nil {
		p = p.next
	}
	p.next = nil

}
func main() {
	head := new(linkedlist)
	var choice string
	for true {
		fmt.Println("enter ur choice")
		fmt.Println("1, insert value in linked list")
		fmt.Println("2, display linked list")
		fmt.Println("3, delete element at first")
		fmt.Println("4, delete element at last")

		fmt.Println("5, Exit")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var data string
			fmt.Println("enter valid value for linked list")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insert(num)
		case "2":
			head.display()
		case "3":
			fmt.Println(head.deletefirst())
		case "4":
			head.deletelast()
		default:
			os.Exit(0)
		}
	}

}

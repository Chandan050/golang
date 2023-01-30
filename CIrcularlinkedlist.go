package main

import "fmt"

type linkedlist struct {
	element int
	next    *linkedlist
}

func main() {
	circle := &linkedlist{}
	circle.insertfirst(10)
	circle.insertfirst(20)
	circle.insertfirst(20)
	circle.insertfirst(20)

	circle.display()

}

func (node *linkedlist) insertfirst(num int) {
	temp := &linkedlist{
		element: num,
		next:    nil,
	}

	if node == nil {
		node = temp
		// lastnode := temp
		return
	}
	temp.next = node

}
func (node *linkedlist) insertlast(num int) {

}

func (node *linkedlist) display() {
	p := node.next
	for p != nil {
		fmt.Println(node.element, "-->")
		p = p.next
	}

}

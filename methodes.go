package main

import (
	"fmt"
	"time"
)

type names []string

// METHODS
func (n names) print() {
	for _, nu := range n {
		fmt.Printf("nu:\t %v\n", nu)
	}
}

type car struct {
	brand string
	price int
}

func main() {

	const day = 24 * time.Hour
	fmt.Printf("type day is %T and time now is %#v", day, day)
	friends := names{"abhi", "ashish", "madhesh", "raghu", "shiva"}

	friends.print()
	fmt.Println(`-------\       /---------------`)
	fmt.Println(`--------\     /----------------`)
	names.print(friends)

	//METHODE WITH POINTER RECEIVER
	mycar := car{brand: "audi", price: 95}
	fmt.Println(mycar)

	fmt.Println(`------------------------------------`)

	changeCar(mycar, "wolks", 20)
	fmt.Println(mycar)

	fmt.Println(`------------------------------------`)

	mycar.changeCar1("tata", 14)
	fmt.Println(mycar)

	fmt.Println(`------------------------------------`)

	(&mycar).changeCar2("mahindra", 24)
	fmt.Println(mycar)

}

// normal func
func changeCar(c car, newCar string, newPrice int) {
	c.brand = newCar
	c.price = newPrice

}

// METHODE BY USAGE
func (c car) changeCar1(newCar string, newPrice int) {
	c.brand = newCar
	c.price = newPrice

}

// POINTER METHODE
func (c *car) changeCar2(newCar string, newPrice int) {
	(*c).brand = newCar
	(*c).price = newPrice

}

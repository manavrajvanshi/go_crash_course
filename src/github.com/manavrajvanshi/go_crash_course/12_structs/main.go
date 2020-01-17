package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	city string
	gender string
	age int
}

func (p Person) greet() string{
	return "Hello, my name is "+p.name+" ,I am "+ strconv.Itoa(p.age)
}

func (p* Person) hasBday () {
	p.age++
	fmt.Println(*p)
}

func main(){
	person1 := Person { name: "Manav", city:"San Jose", gender : "M", age : 22 }
	person1.age = 23
	fmt.Println( person1)

	fmt.Println( person1.greet())
	fmt.Println("Bday")
	person1.hasBday();
	fmt.Println( person1.greet() );
}

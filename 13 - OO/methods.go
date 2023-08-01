package main

import "fmt"

func main() {
	callMethod()
	callUserMethods()
}

// methods
type myInt int // need a type to bind a method

func (i myInt) isEven() bool { //method receiver
	return int(i)%2 == 0
}

func callMethod() {
	mint := myInt(10)
	fmt.Println(mint.isEven())
}

type user struct {
	id   int
	name string
}

func (u user) String() string { // value receiver
	return fmt.Sprintf("%v %v\n", u.name, u.id)
}

func (u *user) UpdateName(name string) { // pointer receiver
	u.name = name
}

func callUserMethods() {
	u := user{
		1,
		"John",
	}
	fmt.Println(u.String())
	u.UpdateName("Paul")
	fmt.Println(u.String())
}

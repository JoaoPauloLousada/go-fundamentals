package main

import "fmt"

func main() {
	var p printer
	p = user{
		"John",
		10,
	}
	fmt.Println(p.Print())

	if u, ok := p.(user); ok != false {
		u.Foo()
	}
}

type printer interface {
	Print() string
}

type user struct {
	name string
	id   int
}

func (u user) Print() string {
	return fmt.Sprintf("%v %v", u.name, u.id)
}

func (u user) Foo() {
	fmt.Println(u.name + " foo")
}

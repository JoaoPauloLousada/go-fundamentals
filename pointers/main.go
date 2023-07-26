package main

import "fmt"

func main() {
	a := 42

	b := &a
	fmt.Println(b, *b)

	a = 27
	fmt.Println(b, *b)

	s := "Hello, world!"
	p := &s
	fmt.Println(p)
	fmt.Println(*p)
	*p = "Hello gophers!"
	fmt.Println(s)

	p = new(string)
	fmt.Println(p, *p)
}

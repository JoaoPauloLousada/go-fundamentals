package main

import "fmt"

func main() {
	a, b := 10, 5
	c := a == b // false - equality
	fmt.Println(c)
	c = a != b // true - inequality
	fmt.Println(c)
	c = a < b // false - less than
	fmt.Println(c)
	c = a <= b // false - less than or equal
	fmt.Println(c)
	c = a > b // true - greater
	fmt.Println(c)
	c = a >= b // true - greater than or equal
	fmt.Println(c)
}

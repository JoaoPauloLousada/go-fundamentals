package main

import "fmt"

func main() {
	// i := 2
	switch i := 2; i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	default:
		fmt.Println("i is greater than 2")
	}
}

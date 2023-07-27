package main

import "fmt"

func main() {
	i := 6
	if i < 5 {
		fmt.Println("i is less than 5")
	} else if i == 5 {
		fmt.Println("i is  5")
	} else {
		fmt.Println("i is greater than 5")
	}
	fmt.Println("After if blocks")
}

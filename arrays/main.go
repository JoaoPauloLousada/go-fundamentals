package main

import (
	"fmt"
)

func main() {
	// arrayDemo()
	slicesDemo()
}

func arrayDemo() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(arr)

	fmt.Println(arr[1])
	arr[1] = "Chai tea"

	fmt.Println(arr)

	arr2 := arr
	fmt.Println(arr2 == arr)
	fmt.Println(&arr2 == &arr)

	arr2[2] = "Chai Latte"
	fmt.Println(arr, arr2)
}

func slicesDemo() {
	var s []string
	fmt.Println(s)

	s = []string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(s)

	s2 := s

	s2[2] = "Chai Latte"

	fmt.Println(s2, s) // share same data

	s = append(s, "Hot Cocolate", "Hot Tea")
	fmt.Println(s, s2)

}

package main

import "fmt"

func main() {
	// success case
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	// error case
	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))
}

func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}

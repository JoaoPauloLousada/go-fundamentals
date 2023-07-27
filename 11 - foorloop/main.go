package main

import "fmt"

func main() {
	//  infinite loop
	// for {...}
	// infiniteLoop()

	// loops until condition
	//for condition {...}
	// loopTillCondition()

	// counter based loop
	// for initializer; test; post clause {...}
	// counterBasedLoop()

	// looping over collecions (arrays, slices and maps)
	// for key, value := range collection {...}
	// for key := range collection {...}
	// for _, value := range collection {...}
	// arrayLoop()
	sliceLoop()
	mapLoop()
}

func infiniteLoop() {
	i := 0
	for {
		fmt.Println(i)
		i += 1
	}
}

func loopTillCondition() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("Done!")
}

func counterBasedLoop() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done!")
}

func arrayLoop() {
	arr := []int{100, 101, 102, 103}
	for key, value := range arr {
		fmt.Println(key, value)
	}
	fmt.Println("Done!")
}

func sliceLoop() {
	println("Slice loop")
	s := []int{10, 20, 30, 40, 50}
	for key, value := range s {
		fmt.Println(key, value)
	}
	fmt.Println("Done!")
}

func mapLoop() {
	println("Map loop")
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println("Done!")

}

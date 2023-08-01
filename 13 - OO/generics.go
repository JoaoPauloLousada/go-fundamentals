package main

import "fmt"

func main() {
	testScores := []float32{
		87.4,
		10.5,
		98, 3,
	}

	c := clone(testScores)
	fmt.Println(testScores, c)
	fmt.Println(&testScores[0], &c[0])

	testScores2 := []float64{
		87.4,
		10.5,
		98, 3,
	}
	c2 := clone(testScores2)
	fmt.Println(testScores2, c2)
	fmt.Println(&testScores2[0], &c2[0])

	testScores3 := map[string]float64{
		"John": 87.4,
		"Doe":  10.5,
		"Jane": 98,
	}

	c3 := cloneMap(testScores3)
	fmt.Println(testScores3, c3)
	john := testScores3["John"]
	johnClone := c3["John"]
	fmt.Println(&john, &johnClone)

	s1 := []int{1, 2, 3, 4}
	fmt.Println(add(s1))
	s2 := []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(add(s2))

}

func clone[T any](s []T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func cloneMap[K comparable, T any](m map[K]T) map[K]T {
	result := make(map[K]T, len(m))
	for i, v := range m {
		result[i] = v
	}
	return result
}

func add[T int | float64](s []T) T {
	var result T
	for _, v := range s {
		result += v
	}
	return result
}

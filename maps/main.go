package main

import "fmt"

func main() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Tea", "Chai Tea"},
	}

	fmt.Println(m)
	fmt.Println(m["coffee"])

	m["other"] = []string{"Hot chocolate"}
	fmt.Println(m) // order doesn't keeps

	delete(m, "tea")
	fmt.Println(m)

	v, ok := m["tea"] // returns default value and boolean flag
	fmt.Println(v, ok)

	m2 := m // share data
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot tea"}
	fmt.Println(m)
	fmt.Println(m2)
}

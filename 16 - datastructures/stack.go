package main

import "fmt"

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop will remove a value at the end
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}

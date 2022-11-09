package main

import "fmt"

type Stack struct {
	data []int
}

// Pop removes the last item from the list and returns it
func (s *Stack) Pop() int {
	last := len(s.data) - 1
	val := s.data[last]
	s.data = s.data[:last]
	return val
}

// push adds a value  to the list
func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func main() {
	s := Stack{}
	s.Push(10)
	s.Push(11)
	s.Push(12)
	s.Push(13)
	fmt.Println(s.data)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

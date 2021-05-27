package main

import "fmt"

func main() {
	var s Stack
	s.push(1)
	s.push(2)
	s.push(3)
	for !s.isEmpty() {
		x, y := s.pop()
		if y {
			fmt.Println(x)
		}
	}
}

type Stack []int

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(element int) {
	*s = append(*s, element)
}

func (s *Stack) pop() (int, bool) {
	if s.isEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

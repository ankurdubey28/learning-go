package main

import "fmt"

type Stack[T comparable] struct {
	val []T
}

func (s *Stack[T]) Push(val T) {
	s.val = append(s.val, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.val) == 0 {
		var zero T
		return zero, false // interesting way of handling zero value with return in generic functions/methods.
	}
	top := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return top, true
}

func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.val {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	var intStack Stack[int]
	intStack.Push(5)
	intStack.Push(9)
	fmt.Println(intStack.Contains(5))
	v, ok := intStack.Pop()
	if !ok {
		fmt.Println("nahi mila")
	} else {
		fmt.Println(v)
	}
}

package utils

import "fmt"

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(data T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) Top() (T, bool) {
	var x T
	if !s.IsEmpty() {
		x = s.items[len(s.items)-1]
		return x, true
	}
	return x, false
}

func (s *Stack[T]) Pop() (T, bool) {
	var x T
	if s.IsEmpty() {
		return x, false
	}
	x, s.items = s.items[len(s.items)-1], s.items[:len(s.items)-1]
	return x, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

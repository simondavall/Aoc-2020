package aoc

import "errors"

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(data T) {
	*s = append((*s), data)
}

func (s *Stack[T]) Pop() (T, error) {
	var popped T
	if s.IsEmpty() {
		return popped, errors.New("Stack is empty")
	} else {
		index := len(*s) - 1
		popped = (*s)[index]
		*s = (*s)[:index]
	}
	return popped, nil
}

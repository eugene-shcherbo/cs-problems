package stack

import (
	"cmp"
	"errors"
)

type Stack[T cmp.Ordered] struct {
	items []T
}

func NewStack[T cmp.Ordered]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) Top() (T, error) {
	if s.Empty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func Sort[T cmp.Ordered](input Stack[T]) Stack[T] {
	var tmp Stack[T]

	for !input.Empty() {
		el, _ := input.Pop()
		tmp.Push(el)
	}

	for !tmp.Empty() {
		toSort, _ := tmp.Pop()

		for !input.Empty() {
			el, _ := input.Top()
			if el, _ := input.Top(); el <= toSort {
				break
			}
			input.Pop()
			tmp.Push(el)
		}

		input.Push(toSort)
	}

	return input
}

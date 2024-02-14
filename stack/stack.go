package stack

import (
	"errors"
	"fmt"
	"sync"
)

type Stack[T comparable] struct {
	top    int
	mu     *sync.RWMutex
	length int
	stack  []T
}

func NewStack[T comparable](length int) *Stack[T] {
	s := &Stack[T]{
		top:    -1,
		mu:     &sync.RWMutex{},
		length: length,
		stack:  make([]T, 0, length),
	}
	return s
}

func (s *Stack[T]) AddToStack(element T) error {
	// handle overflow
	if s.top == s.length {
		return errors.New("overflow")
	}
	s.top += 1
	s.stack = append(s.stack, element)
	return nil
}

func (s *Stack[T]) RemoveFromStack() (T, error) {
	var newValue T

	// handle underflow
	if s.top == -1 {
		return newValue, errors.New("underflow")
	}
	s.top -= 1
	topElement := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return topElement, nil
}

func (s *Stack[T]) PrintStack() {
	fmt.Println("Printing Stack...")
	for _, v := range s.stack {
		fmt.Printf("%v ", v)
	}
	fmt.Println("")
}

func (s *Stack[T]) Length() int {
	return s.top
}

func (s *Stack[T]) Top() T {

	var newValue T

	if s.top == -1 {
		return newValue
	}

	topElement := s.stack[s.top]
	return topElement
}

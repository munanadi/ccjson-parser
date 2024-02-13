package stack

import (
	"errors"
	"fmt"
	"sync"
)

type Stack struct {
	top    int
	mu     *sync.RWMutex
	length int
	stack  []byte
}

func NewStack(length int) *Stack {
	s := &Stack{
		top:    -1,
		mu:     &sync.RWMutex{},
		length: length,
		stack:  make([]byte, 0, length),
	}
	return s
}

func (s *Stack) AddToStack(element byte) error {
	// handle overflow
	if s.top == s.length {
		return errors.New("overflow")
	}
	s.top += 1
	s.stack = append(s.stack, element)
	return nil
}

func (s *Stack) RemoveFromStack() (byte, error) {
	// handle underflow
	if s.top == -1 {
		return byte(0), errors.New("underflow")
	}
	s.top -= 1
	topElement := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return topElement, nil
}

func (s *Stack) PrintStack() {
	fmt.Println("Printing Stack...")
	fmt.Println(s.stack)
	fmt.Println("")
}

func (s *Stack) Length() int {
	return s.top
}

func (s *Stack) Top() byte {

	if s.top == -1 {
		return byte(0)
	}

	topElement := s.stack[s.top]
	return topElement
}

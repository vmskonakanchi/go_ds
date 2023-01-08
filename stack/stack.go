package stack

import (
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type Stack[T comparable] struct {
	count uint
	head  *Node[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		head: nil,
	}
}

func (s *Stack[T]) Push(item T) bool {
	if s.head == nil {
		s.head = &Node[T]{
			data: item,
			next: nil,
		}
	} else {
		n := &Node[T]{
			data: item,
			next: nil,
		}
		n.next = s.head
		s.head = n
	}
	s.count++
	return true
}

func (s *Stack[T]) Pop() bool {
	if s.count > 0 {
		if s.head.next == nil {
			s.head = nil
		} else {
			s.head = s.head.next
		}
		s.count--
		return true
	}
	fmt.Println("There are no elements to pop")
	return false
}

func (s *Stack[T]) IsEmpty() bool {
	return s.count == 0
}

func (s *Stack[T]) Peek() T {
	var result T
	if s.count > 0 {
		result = s.head.data
	}
	return result
}

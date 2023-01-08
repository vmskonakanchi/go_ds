package queue

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type Queue[T comparable] struct {
	count uint
	head  *Node[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		count: 0,
		head:  nil,
	}
}

func (q *Queue[T]) Enqueue(item T) bool {
	if q.head == nil {
		q.head = &Node[T]{
			data: item,
			next: nil,
		}
		q.count++
		return true
	}
	temp := q.head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &Node[T]{
		data: item,
		next: nil,
	}
	q.count++
	return true
}

func (q *Queue[T]) Dequeue() bool {
	if q.count > 0 {
		if q.head.next == nil {
			q.head = nil
		} else {
			q.head = q.head.next
		}
		q.count--
		return true
	}
	fmt.Println("There are no elements to dequeue")
	return false
}

func (q *Queue[T]) Peek() T {
	var result T
	if q.count > 0 {
		result = q.head.data
	}
	return result
}

func (q *Queue[T]) IsEmpty() bool {
	return q.count == 0
}

package linkedlist

import (
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	size uint
	head *Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	llist := &LinkedList[T]{
		size: 0,
		head: nil,
	}
	return llist
}

/* prints in the linked list formatted way */
func (list *LinkedList[T]) Print() {
	if list.size == 0 {
		return
	}
	temp := list.head
	fmt.Printf("%v", temp.data)
	for temp.next != nil {
		fmt.Printf(" -> %v", temp.next.data)
		temp = temp.next
	}
}

/*
Adds the given data at the head of the linked list

returns : true if the insertion is successfull else false
*/
func (list *LinkedList[T]) InsertAtHead(item T) bool {
	if list.head == nil {
		list.head = &Node[T]{
			data: item,
			next: nil,
		}
		list.size++
		return true
	}
	temp := list.head
	n := &Node[T]{
		data: item,
		next: temp,
	}
	list.head = n
	list.size++
	return true
}

/*
Adds the given data at the end of the linked list

returns : true if the insertion is successfull else false
*/
func (list *LinkedList[T]) InsertAtEnd(item T) bool {
	if list.head == nil {
		list.head = &Node[T]{
			data: item,
			next: nil,
		}
		list.size++
		return true
	}
	temp := list.head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &Node[T]{
		data: item,
		next: nil,
	}
	list.size++
	return true
}

func (list *LinkedList[T]) Remove(position uint) bool {
	if list.size == 0 {
		fmt.Println("There are no elements in the list to search")
		return false
	}
	if position > list.size {
		fmt.Println("The position is not valid")
		return false
	}
	if position == 0 {
		if list.head != nil {
			if list.head.next != nil {
				list.head = list.head.next
			} else {
				list.head = nil
			}
			list.size--
			return true
		}
		fmt.Println("There is no head elements to delete")
		return false
	}
	prev := list.head
	curr := list.head
	for i := 0; i < int(position); i++ {
		prev = curr
		curr = curr.next
	}
	if curr.next != nil {
		prev.next = curr.next
	} else {
		prev.next = nil
	}
	list.size--
	return true
}

func (list *LinkedList[T]) Contains(item T) bool {
	if list.head == nil {
		return false
	}
	if list.size == 1 {
		if list.head.data == item {
			return true
		}
	}

	temp := list.head

	for temp.next != nil {
		if temp.data == item {
			return true
		}
		temp = temp.next
	}
	return false
}

// func (list *LinkedList[T]) Sort() {
// 	//TODO: Sort the elements in the list by exchanging values
// }

package main

import (
	"fmt"

	"github.com/vmskonakanchi/go_ds/linkedlist"
)

func main() {
	list := linkedlist.NewLinkedList[int]()
	list.InsertAtEnd(0)
	list.InsertAtEnd(2)
	list.InsertAtEnd(4)
	list.Print()
	fmt.Println()
}

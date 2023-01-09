package main

import (
	"fmt"

	"github.com/vmskonakanchi/go_ds/tree"
)

func main() {
	tree := tree.NewBinaryTree[int]()

	for i := 0; i < 6; i++ {
		tree.Insert(i + 1)
	}

	tree.Traverse("inorder")

	fmt.Println("the number contains ", tree.Contains(2))
}

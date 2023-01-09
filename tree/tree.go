package tree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	data        T
	left, right *Node[T]
}

type Tree[T constraints.Ordered] struct {
	count uint
	root  *Node[T]
}

func NewBinaryTree[T constraints.Ordered]() *Tree[T] {
	return &Tree[T]{
		root: nil,
	}
}

func print[T constraints.Ordered](method string, root *Node[T]) {
	if root != nil {
		switch method {
		case "inorder":
			print(method, root.left)
			fmt.Println(root.data)
			print(method, root.right)
		case "postorder":
			print(method, root.left)
			print(method, root.right)
			fmt.Println(root.data)
		case "preorder":
			fmt.Println(root.data)
			print(method, root.left)
			print(method, root.right)
		default:
			fmt.Println("Travsering failed Please enter only these `preorder` , `postorder` , `inorder`")
		}

	}
}

func (t *Tree[T]) Traverse(method string) {
	method = strings.ToLower(method)
	print(method, t.root)
}

// helper function to insert data into binary search tree
func addToTree[T constraints.Ordered](root *Node[T], item T) {
	if item < root.data {
		if root.left == nil {
			root.left = &Node[T]{
				data:  item,
				left:  nil,
				right: nil,
			}
		} else {
			addToTree(root.left, item)
		}
	} else {
		if root.right == nil {
			root.right = &Node[T]{
				data:  item,
				left:  nil,
				right: nil,
			}
		} else {
			addToTree(root.right, item)
		}
	}
}

func (t *Tree[T]) Insert(item T) bool {
	if t.root == nil {
		t.root = &Node[T]{
			data:  item,
			left:  nil,
			right: nil,
		}
	} else {
		addToTree(t.root, item)
	}
	t.count++
	return true
}

func find[T constraints.Ordered](item T, node *Node[T]) bool {
	if item == node.data {
		return true
	}
	if item < node.data {
		if node.left == nil {
			return false
		} else {
			return find(item, node.left)
		}
	} else {
		if node.right == nil {
			return false
		} else {
			return find(item, node.right)
		}
	}
}

func (t *Tree[T]) Contains(item T) bool {
	return find(item, t.root)
}

// func (t *Tree[T]) Remove(item T) bool {
// }

func (t *Tree[T]) Size() uint {
	return t.count
}

func (t *Tree[T]) IsEmpty() bool {
	return t.count == 0
}

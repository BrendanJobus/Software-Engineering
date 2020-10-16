package main

// Notes: to make changes to an array or slice that persist after a functions scope, use a pointer to that slice

import (
	"errors"
	"fmt"
	"log"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {		//
	values := []int{10, 5, 13, 6, 4, 15, 2, 7, 35}							//						10
																															//				  5	  13
																															//		    4  6	 15
	tree := &Tree{}																							//		   2    7 	 35
	for i := 0; i < len(values); i++ {													//
		err := tree.Insert(values[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}
	x := findLCA(tree.Root, 2, 7)
	fmt.Println(x)
}

func findPath(root *Node, path *[]int,k int) bool {
	if (root == nil) { return false }

	*path = append(*path, root.Value)

	if root.Value == k { return true }

	if (root.Left != nil && findPath(root.Left, path, k)) || (root.Right != nil && findPath(root.Right, path, k)) { return true }

	*path = (*path)[:len(*path) - 1]
	return false
}

func findLCA(root *Node, n1, n2 int) int {
	var path1, path2 []int

	if !findPath(root, &path1, n1) || !findPath(root, &path2, n2) {
		return -1
	}

	var i int
	for i = 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] != path2[i] {
			break
		}
	}
	if i > 0 {
		return path1[i - 1]
	}
	return path1[i]
}

func (n *Node) Insert(value int) error {

	if n == nil {
		return errors.New("Cannot insert a value into nil tree")
	}

	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return nil
		}
		return n.Left.Insert(value)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value:value}
			return nil
		}
		return n.Right.Insert(value)
	}
	return nil
}

func (n *Node) Find(s int) (int, bool) {
	if n == nil {
		return -1, false
	}

	switch {
	case s == n.Value:
		return n.Value, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (t *Tree) Insert(value int) error {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return nil
	}
	return t.Root.Insert(value)
}

func (t *Tree) Find(s int) (int, bool) {
	if t.Root == nil {
		return -1, false
	}
	return t.Root.Find(s)
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

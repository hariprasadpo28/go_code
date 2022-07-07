package main

import (
	"fmt"
)

type tree struct {
	left *tree
	right *tree
	value string
}

func inorder(t *tree) {
	if t != nil{
		inorder(t.left)
		fmt.Println(t.value)
		inorder(t.right)
	}
}

func preorder(t *tree)  {

	if t != nil {
		fmt.Println(t.value)
		preorder(t.left)
		preorder(t.right)

	}
}

func postorder(t *tree)  {
	if t == nil{
		return
	}
	postorder(t.left)
	postorder(t.right)
	fmt.Println(t.value)
}

func main() {

	var head = tree{nil,nil,"+"}
	head.left = &tree{nil,nil,"a"}
	head.right = &tree{nil,nil,"-"}
	head.right.left = &tree{nil,nil,"b"}
	head.right.right = &tree{nil,nil,"c"}

	fmt.Println("Pre-order Traversal\n")
	preorder(&head)
	fmt.Println("\n")

	fmt.Println("Inorder Traversal\n")
	inorder(&head)
	fmt.Println("\n")

	fmt.Println("Postorder Traversal\n")
	postorder(&head)

}

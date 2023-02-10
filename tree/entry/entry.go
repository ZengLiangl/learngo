package main

import (
	"fmt"
	"learngo/tree"

	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	node := myTreeNode{myNode.node.Left}
	node.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(100)
	s.Insert(10000)
	s.Insert(1000000)
	fmt.Println(s.Has(100))
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.Setval(4)

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount: ", nodeCount)
	//node := myTreeNode{&root}
	//node.postOrder()
	// testSparse()

	maxNode := 0
	c := root.TraverseWithChannel()
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("MaxVal :", maxNode)
}

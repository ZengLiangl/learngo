package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) Setval(val int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = val
}
func CreateNode(val int) *Node {
	// 相当于构造方法
	return &Node{Value: val}
}

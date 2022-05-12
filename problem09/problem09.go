package problem09

type Node struct {
	Val      int
	Children *[]Node
}

func NewNode(val int, childrens *[]Node) *Node {
	return &Node{
		val,
		childrens,
	}
}

func SumTree(root *Node, sum int) int {
	if root == nil {
		return 0
	}

	for _, child := range *root.Children {
		sum = SumTree(&child, sum)
	}
	return sum + root.Val
}

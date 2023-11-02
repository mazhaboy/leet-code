package tree

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		Invert(root)
	}
	return root
}

func Invert(node *TreeNode) {
	if node == nil {
		return
	}
	var left, right *TreeNode
	if node.Left != nil {
		left = node.Left
	}
	if node.Right != nil {
		right = node.Right
	}
	node.Right = left
	node.Left = right
	Invert(node.Left)
	Invert(node.Right)
}

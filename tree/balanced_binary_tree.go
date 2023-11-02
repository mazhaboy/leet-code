package tree

func isBalanced(root *TreeNode) bool {
	arr := []bool{true}
	balance(root, arr)
	return arr[0]
}

func balance(node *TreeNode, arr []bool) int {
	if node == nil {
		return 0
	}

	left := balance(node.Left, arr)
	right := balance(node.Right, arr)

	if left-right > 1 || left-right < -1 {
		arr[0] = false
	}

	return 1 + Max(left, right)
}

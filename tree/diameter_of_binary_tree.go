package tree

func DiameterOfBinaryTree(root *TreeNode) int {
	arr := []int{0}
	dfs(root, arr)
	return arr[0]
}

func dfs(node *TreeNode, max []int) int {
	if node == nil {
		return 0
	}

	l := dfs(node.Left, max)
	r := dfs(node.Right, max)

	max[0] = Max(max[0], r+l)
	return 1 + Max(r, l)
}

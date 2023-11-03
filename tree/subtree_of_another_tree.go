package tree

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if compare(root, subRoot) {
		return true
	}

	if IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot) {
		return true
	}

	return false
}

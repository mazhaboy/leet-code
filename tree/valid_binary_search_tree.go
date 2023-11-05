package tree

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(node *TreeNode, leftVal, rightVal int) bool {
	if node == nil {
		return true
	}
	if !(leftVal < node.Val && node.Val < rightVal) {
		return false
	}

	if !check(node.Left, leftVal, node.Val) {
		return false
	}
	if !check(node.Right, node.Val, rightVal) {
		return false
	}

	return true
}

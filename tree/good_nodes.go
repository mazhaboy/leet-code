package tree

import (
	"fmt"
)

func GoodNodes(root *TreeNode) int {
	return countGoodNodes(root, root.Val)
}

func countGoodNodes(node *TreeNode, max int) int {
	counter := 0
	if node == nil {
		return 0
	}

	max = Max(max, node.Val)
	if node.Val >= max {
		fmt.Println(max)
		counter++
	}

	counter += countGoodNodes(node.Left, max)
	counter += countGoodNodes(node.Right, max)

	return counter
}

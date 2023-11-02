package tree

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := Constructor()
	q.PushBack(root)
	level := 0
	for !q.Empty() {
		for range q.Arr {
			node, _ := q.PopFront()
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		level++
	}
	return level
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + Max(depth(node.Left), depth(node.Right))
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

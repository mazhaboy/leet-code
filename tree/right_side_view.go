package tree

func rightSideView(root *TreeNode) []int {
	arr := make([]int, 0)
	if root == nil {
		return arr
	}

	queue := Constructor()
	queue.PushFront(root)
	for !queue.Empty() {
		level := make([]int, 0)
		for range queue.Arr {
			node, _ := queue.PopBack()
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			level = append(level, node.Val)
		}
		if len(level) != 1 {
			arr = append(arr, level[len(level)-1])
		} else {
			arr = append(arr, level[0])
		}
	}

	return arr
}

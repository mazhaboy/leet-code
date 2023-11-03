package tree

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return compare(p, q)
}

func compare(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if !compare(p.Left, q.Left) || !compare(p.Right, q.Right) {
			return false
		}
	}

	return true
}

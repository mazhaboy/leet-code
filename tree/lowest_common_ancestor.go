package tree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	lca := &TreeNode{}

	if (root.Val > p.Val && root.Val < q.Val) || (root.Val < p.Val && root.Val > q.Val) || (root.Val == q.Val) || (root.Val == p.Val) {
		return root
	} else if root.Val > p.Val && root.Val > q.Val {
		lca = LowestCommonAncestor(root.Left, q, p)
	} else if root.Val < p.Val && root.Val < q.Val {
		lca = LowestCommonAncestor(root.Right, q, p)
	}

	return lca
}

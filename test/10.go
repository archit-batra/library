package main

// TreeNode definition (same as maxDepthBinaryTree.go)

// lowestCommonAncestor finds LCA in BST.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root
	for current != nil {
		if p.Val < current.Val && q.Val < current.Val {
			current = current.Left // Both left
		} else if p.Val > current.Val && q.Val > current.Val {
			current = current.Right // Both right
		} else {
			return current // Split or one is current
		}
	}
	return nil // Should not happen
}

// Helper to build tree (similar to before, but for BST assume input is level order)

// Find node by value (for p and q)
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return findNode(root.Left, val)
	}
	return findNode(root.Right, val)
}

package main

// TreeNode definition (same as maxDepthBinaryTree.go)

// isBalanced checks if the tree is height-balanced.
func isBalanced(root *TreeNode) bool {
	var getHeight func(node *TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getHeight(node.Left)
		right := getHeight(node.Right)
		if left == -1 || right == -1 || abs(left-right) > 1 {
			return -1 // Unbalanced
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	return getHeight(root) != -1
}

// abs helper
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Helper to build tree (similar to before)

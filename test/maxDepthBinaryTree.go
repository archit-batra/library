package main

// TreeNode definition.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth computes the maximum depth of the binary tree.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0 // Base case
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// Helper to build tree (level order, null as nil)
func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		current := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			current.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, current.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			current.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}

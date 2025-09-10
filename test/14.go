package main

import "fmt"

// TreeNode definition (same as maxDepthBinaryTree.go)

// invertTree inverts the binary tree.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left // Swap
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// buildTree from earlier

// printTree level order
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*TreeNode{root}
	vals := []int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		vals = append(vals, current.Val)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	fmt.Println(vals)
}

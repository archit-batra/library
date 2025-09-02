package main

// ListNode definition in mergeTwoLists.go

// hasCycle detects cycle using Floyd's algorithm.
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Slow: 1 step
		fast = fast.Next.Next // Fast: 2 steps
		if slow == fast {
			return true // Cycle detected
		}
	}
	return false
}

// Helper to build list with cycle (pos = -1 for no cycle)
func buildListWithCycle(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	var cycleNode *ListNode
	for i, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
		if i+1 == pos {
			cycleNode = current
		}
	}
	if pos >= 0 {
		current.Next = cycleNode
	}
	return head
}

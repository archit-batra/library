package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy head
	current := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			current.Next = p1
			p1 = p1.Next
		} else {
			current.Next = p2
			p2 = p2.Next
		}
		current = current.Next
	}
	if p1 != nil {
		current.Next = p1 // Append remaining
	} else {
		current.Next = p2
	}
	return dummy.Next
}

// Helper to build list from slice
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper to print list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

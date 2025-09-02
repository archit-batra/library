package main

import "fmt"

func main() {
	// --- twoSum.go: Testing twoSum ---
	fmt.Println("[twoSum.go] Testing twoSum")
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Printf("Input: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("Output: %v\n\n", twoSum(nums1, target1))

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Printf("Input: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %v\n\n", twoSum(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Printf("Input: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("Output: %v\n\n", twoSum(nums3, target3))

	// --- validParentheses.go: Testing isValid ---
	fmt.Println("[validParentheses.go] Testing isValid")
	s1 := "()"
	fmt.Printf("Input: s = \"%s\"\n", s1)
	fmt.Printf("Output: %t\n\n", isValid(s1))

	s2 := "()[]{}"
	fmt.Printf("Input: s = \"%s\"\n", s2)
	fmt.Printf("Output: %t\n\n", isValid(s2))

	s3 := "(]"
	fmt.Printf("Input: s = \"%s\"\n", s3)
	fmt.Printf("Output: %t\n\n", isValid(s3))

	// --- mergeTwoLists.go: Testing mergeTwoLists ---
	fmt.Println("[mergeTwoLists.go] Testing mergeTwoLists")
	list1_1 := buildList([]int{1, 2, 4})
	list2_1 := buildList([]int{1, 3, 4})
	fmt.Print("Input: list1 = [1,2,4], list2 = [1,3,4]\nOutput: ")
	printList(mergeTwoLists(list1_1, list2_1))

	list1_2 := buildList([]int{})
	list2_2 := buildList([]int{})
	fmt.Print("Input: list1 = [], list2 = []\nOutput: ")
	printList(mergeTwoLists(list1_2, list2_2))

	list1_3 := buildList([]int{})
	list2_3 := buildList([]int{0})
	fmt.Print("Input: list1 = [], list2 = [0]\nOutput: ")
	printList(mergeTwoLists(list1_3, list2_3))

	// --- maxProfit.go: Testing maxProfit ---
	fmt.Println("[maxProfit.go] Testing maxProfit")
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Input: prices = %v\n", prices1)
	fmt.Printf("Output: %d\n\n", maxProfit(prices1))

	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Printf("Input: prices = %v\n", prices2)
	fmt.Printf("Output: %d\n\n", maxProfit(prices2))

	// --- isPalindrome.go: Testing isPalindrome ---
	fmt.Println("[isPalindrome.go] Testing isPalindrome")
	p1 := "A man, a plan, a canal: Panama"
	fmt.Printf("Input: s = \"%s\"\n", p1)
	fmt.Printf("Output: %t\n\n", isPalindrome(p1))

	p2 := "race a car"
	fmt.Printf("Input: s = \"%s\"\n", p2)
	fmt.Printf("Output: %t\n\n", isPalindrome(p2))

	p3 := " "
	fmt.Printf("Input: s = \"%s\"\n", p3)
	fmt.Printf("Output: %t\n\n", isPalindrome(p3))

	// --- hasCycle.go: Testing hasCycle ---
	fmt.Println("[hasCycle.go] Testing hasCycle")
	head1 := buildListWithCycle([]int{3, 2, 0, -4}, 1)
	fmt.Println("Input: [3,2,0,-4] with cycle at pos 1")
	fmt.Printf("Output: %t\n\n", hasCycle(head1))

	head2 := buildListWithCycle([]int{1, 2}, 0)
	fmt.Println("Input: [1,2] with cycle at pos 0")
	fmt.Printf("Output: %t\n\n", hasCycle(head2))

	head3 := buildListWithCycle([]int{1}, -1)
	fmt.Println("Input: [1] with no cycle")
	fmt.Printf("Output: %t\n\n", hasCycle(head3))

	// --- search.go: Testing search ---
	fmt.Println("[search.go] Testing search")
	nums_1 := []int{-1, 0, 3, 5, 9, 12}
	target_1 := 9
	fmt.Printf("Input: nums = %v, target = %d\n", nums_1, target_1)
	fmt.Printf("Output: %d\n\n", search(nums_1, target_1))

	target_2 := 2
	fmt.Printf("Input: nums = %v, target = %d\n", nums_1, target_2)
	fmt.Printf("Output: %d\n\n", search(nums_1, target_2))

	// --- floodFill.go: Testing floodFill ---
	fmt.Println("[floodFill.go] Testing floodFill")
	image1 := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr1, sc1, color1 := 1, 1, 2
	fmt.Printf("Input: image = %v, sr = %d, sc = %d, color = %d\n", image1, sr1, sc1, color1)
	fmt.Println("Output:")
	printImage(floodFill(image1, sr1, sc1, color1))
	fmt.Println()

	image2 := [][]int{{0, 0, 0}, {0, 0, 0}}
	sr2, sc2, color2 := 0, 0, 0
	fmt.Printf("Input: image = %v, sr = %d, sc = %d, color = %d\n", image2, sr2, sc2, color2)
	fmt.Println("Output:")
	printImage(floodFill(image2, sr2, sc2, color2))

	// --- lowestCommonAncestor.go: Testing lowestCommonAncestor ---
	fmt.Println("[lowestCommonAncestor.go] Testing lowestCommonAncestor")
	vals1 := []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
	root1 := buildTree(vals1)
	p_1 := findNode(root1, 2)
	q1 := findNode(root1, 8)
	lca1 := lowestCommonAncestor(root1, p_1, q1)
	fmt.Println("Input: root = [6,2,8,0,4,7,9,null,null,3,5], p=2, q=8")
	fmt.Printf("Output: %d\n\n", lca1.Val)

	p_2 := findNode(root1, 2)
	q2 := findNode(root1, 4)
	lca2 := lowestCommonAncestor(root1, p_2, q2)
	fmt.Println("Input: root = [6,2,8,0,4,7,9,null,null,3,5], p=2, q=4")
	fmt.Printf("Output: %d\n\n", lca2.Val)

	// --- isBalanced.go: Testing isBalanced ---
	fmt.Println("[isBalanced.go] Testing isBalanced")
	vals_1 := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root_1 := buildTree(vals_1)
	fmt.Println("Input: [3,9,20,null,null,15,7]")
	fmt.Printf("Output: %t\n\n", isBalanced(root_1))

	vals2 := []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}
	root2 := buildTree(vals2)
	fmt.Println("Input: [1,2,2,3,3,null,null,4,4]")
	fmt.Printf("Output: %t\n\n", isBalanced(root2))

	root3 := buildTree([]interface{}{})
	fmt.Println("Input: []")
	fmt.Printf("Output: %t\n\n", isBalanced(root3))

	// --- queueUsingStack.go: Testing MyQueue ---
	fmt.Println("[queueUsingStack.go] Testing MyQueue")
	q := NewMyQueue()
	q.Push(1)
	q.Push(2)
	fmt.Printf("Peek: %d\n", q.Peek())   // 1
	fmt.Printf("Pop: %d\n", q.Pop())     // 1
	fmt.Printf("Empty: %t\n", q.Empty()) // false

	// --- minStack.go: Testing MinStack ---
	fmt.Println("[minStack.go] Testing MinStack")
	ms := NewMinStack()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Printf("GetMin: %d\n", ms.GetMin()) // -3
	ms.Pop()
	fmt.Printf("Top: %d\n", ms.Top())       // 0
	fmt.Printf("GetMin: %d\n", ms.GetMin()) // -2

	// --- invertTree.go: Testing invertTree ---
	fmt.Println("[invertTree.go] Testing invertTree")
	vals__1 := []interface{}{4, 2, 7, 1, 3, 6, 9}
	root__1 := buildTree(vals__1)
	fmt.Println("Input: [4,2,7,1,3,6,9]")
	inverted1 := invertTree(root__1)
	fmt.Print("Output: ")
	printTree(inverted1)
	fmt.Println()

	vals__2 := []interface{}{2, 1, 3}
	root__2 := buildTree(vals__2)
	fmt.Println("Input: [2,1,3]")
	inverted2 := invertTree(root__2)
	fmt.Print("Output: ")
	printTree(inverted2)
	fmt.Println()

	root__3 := buildTree([]interface{}{})
	fmt.Println("Input: []")
	inverted3 := invertTree(root__3)
	fmt.Print("Output: ")
	printTree(inverted3)
}

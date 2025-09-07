package main

func twoSum(nums []int, target int) []int {
	// Map from value to its index in the array
	idx := make(map[int]int)
	for i, v := range nums {
		// If we have seen target - v before, we found the pair
		if j, ok := idx[target-v]; ok {
			return []int{j, i}
		}
		// Otherwise, remember this value's index
		idx[v] = i
	}
	return nil // Per problem, solution exists; this is a fallback
}

package main

// twoSum returns the indices of the two numbers in 'nums' that add up to 'target'.
// This is a classic "hash map" problem in DSA (Data Structures & Algorithms).
// The function uses a map to remember numbers we've already seen as we loop through the array.
// For each number, we check if we've already seen its "complement" (target - num).
// If yes, we return the indices. If not, we store the current number and its index in the map.
//
// Go syntax notes:
// - 'make(map[int]int)' creates a map from int to int.
// - 'for i, num := range nums' loops through nums, giving index 'i' and value 'num'.
// - 'idx, ok := m[complement]' looks up 'complement' in the map. 'ok' is true if found.
// - 'return []int{idx, i}' returns a slice (dynamic array) of two indices.
//
// Time complexity: O(n) because each number is processed once.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // Create an empty map to store number -> index
	for i, num := range nums {
		complement := target - num // The number we need to find
		// Check if we've already seen the complement
		if idx, ok := m[complement]; ok {
			// If found, return the indices (complement's index and current index)
			return []int{idx, i}
		}
		// Otherwise, remember this number and its index for future checks
		m[num] = i
	}
	// If no solution is found, return nil (problem guarantees a solution exists)
	return nil
}

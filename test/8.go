package main

// search performs binary search on sorted nums for target.
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2 // Avoid overflow
		if nums[mid] == target {
			return mid // Found
		} else if nums[mid] < target {
			low = mid + 1 // Search right
		} else {
			high = mid - 1 // Search left
		}
	}
	return -1 // Not found
}

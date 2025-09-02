package main

import "fmt"

// floodFill performs flood fill starting from (sr, sc) with new color.
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image // No change needed
	}
	rows, cols := len(image), len(image[0])
	oldColor := image[sr][sc]
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || image[r][c] != oldColor {
			return // Out of bounds or wrong color
		}
		image[r][c] = color // Change color
		dfs(r-1, c)         // Up
		dfs(r+1, c)         // Down
		dfs(r, c-1)         // Left
		dfs(r, c+1)         // Right
	}
	dfs(sr, sc)
	return image
}

// Helper to print 2D array
func printImage(image [][]int) {
	for _, row := range image {
		fmt.Println(row)
	}
}

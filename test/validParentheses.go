package main

// isValid checks if the parentheses in s are valid using a stack.
func isValid(s string) bool {
	stack := []rune{}   // Stack to hold opening brackets
	m := map[rune]rune{ // Mapping closing to opening
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char) // Push opening
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[char] {
				return false // Mismatch or empty stack
			}
			stack = stack[:len(stack)-1] // Pop
		}
	}
	return len(stack) == 0 // Valid if stack is empty
}

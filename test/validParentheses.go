package main

import "fmt"

// ValidParentheses checks if the string has valid matching parentheses, brackets, and braces.
// DSA Concept: We use a stack (Last In, First Out - LIFO) to keep track of opening brackets. For each closing bracket, we pop the stack and check if it matches the expected opening. If mismatch or stack not empty at end, invalid. This is O(n) time and space.
// Example: For s = "()[]{}":
// - Push '(', stack=['(']
// - See ')', matches pop '(', stack=[]
// - Push '[', stack=['[']
// - See ']', matches pop '[', stack=[]
// - Push '{', stack=['{']
// - See '}', matches pop '{', stack=[] -> valid
// For s = "(]":
// - Push '(', stack=['(']
// - See ']', doesn't match '(', return false
// This ensures proper nesting and matching.
// Go Syntax: We use a slice []rune as a stack (rune for Unicode chars). append() to push, stack[len-1] to peek, stack = stack[:len-1] to pop. Maps like map[rune]rune for closing-to-opening pairs. range s iterates over characters as runes.

func isValid(s string) bool {
    stack := []rune{} // Initialize an empty slice to act as a stack for opening brackets.
    bracketMap := map[rune]rune{ // Map of closing brackets to their corresponding opening ones.
        ')': '(', // ')' expects '('
        ']': '[', // ']' expects '['
        '}': '{', // '}' expects '{'
    }
    for _, char := range s { // Loop through each character in the string.
        if char == '(' || char == '[' || char == '{' { // If it's an opening bracket.
            stack = append(stack, char) // Push it onto the stack.
        } else if opening, ok := bracketMap[char]; ok { // If it's a closing bracket (exists in map).
            if len(stack) == 0 || stack[len(stack)-1] != opening { // Check if stack is empty or top doesn't match expected opening.
                return false // Mismatch or unexpected closing, invalid.
            }
            stack = stack[:len(stack)-1] // Pop the matching opening from the stack.
        } // Ignore non-bracket characters, as per problem (though assumes only brackets).
    }
    return len(stack) == 0 // If stack is empty at the end, all brackets matched; else unpaired openings left.
}

func main() {
    s := "()" // Example input.
    fmt.Println(isValid(s)) // Output: true
    s = "(]" // Another example.
    fmt.Println(isValid(s)) // Output: false
}

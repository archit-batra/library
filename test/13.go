package main

// MinStack supports min in O(1).
type MinStack struct {
	stack []int // Values
	min   []int // Mins
}

func NewMinStack() MinStack {
	return MinStack{}
}

// push adds val.
func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.min) == 0 || val <= s.min[len(s.min)-1] {
		s.min = append(s.min, val) // New min
	} else {
		s.min = append(s.min, s.min[len(s.min)-1]) // Current min
	}
}

// pop removes top.
func (s *MinStack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
		s.min = s.min[:len(s.min)-1]
	}
}

// top returns top.
func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return -1 // Invalid
	}
	return s.stack[len(s.stack)-1]
}

// getMin returns min.
func (s *MinStack) GetMin() int {
	if len(s.min) == 0 {
		return -1 // Invalid
	}
	return s.min[len(s.min)-1]
}

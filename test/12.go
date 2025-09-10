package main

// MyQueue implements queue using two stacks.
type MyQueue struct {
	input  []int // For push
	output []int // For pop/peek
}

func NewMyQueue() MyQueue {
	return MyQueue{}
}

// push adds x to back.
func (q *MyQueue) Push(x int) {
	q.input = append(q.input, x)
}

// pop removes front and returns it.
func (q *MyQueue) Pop() int {
	q.Peek() // Ensure output has front
	if len(q.output) == 0 {
		return -1 // Invalid, but per constraints valid
	}
	front := q.output[len(q.output)-1]
	q.output = q.output[:len(q.output)-1]
	return front
}

// peek returns front.
func (q *MyQueue) Peek() int {
	if len(q.output) == 0 {
		for len(q.input) > 0 {
			q.output = append(q.output, q.input[len(q.input)-1])
			q.input = q.input[:len(q.input)-1]
		}
	}
	if len(q.output) == 0 {
		return -1 // Empty
	}
	return q.output[len(q.output)-1]
}

// empty checks if queue is empty.
func (q *MyQueue) Empty() bool {
	return len(q.input) == 0 && len(q.output) == 0
}

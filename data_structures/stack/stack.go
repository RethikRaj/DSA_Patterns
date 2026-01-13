package main

import "fmt"

type Stack struct {
	elements    []int
	maxCapacity int
}

func NewStack(maxCapacity int) *Stack {
	return &Stack{
		// Pre-allocate memory for performance, but start with length 0
		elements:    make([]int, 0, maxCapacity),
		maxCapacity: maxCapacity,
	}
}

func (s *Stack) Push(val int) error {
	if len(s.elements) >= s.maxCapacity {
		return fmt.Errorf("stack overflow: capacity %d reached", s.maxCapacity)
	}
	s.elements = append(s.elements, val)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack underflow: stack is empty")
	}

	// Get the last element
	index := len(s.elements) - 1
	val := s.elements[index]

	// Remove the last element from the slice
	s.elements = s.elements[:index]
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func main() {
	myStack := NewStack(3)

	// Push elements
	elementsToPush := []int{10, 20, 30, 40}
	for _, val := range elementsToPush {
		fmt.Printf("Pushing: %d\n", val)
		err := myStack.Push(val)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	topVal, _ := myStack.Peek()
	fmt.Printf("\nPeeking top element: %d\n", topVal)

	// Pop elements until empty
	fmt.Println("\n--- Popping elements ---")
	for i := 0; i < 4; i++ { // Attempting 4 pops on a 3-element stack
		val, err := myStack.Pop()
		if err != nil {
			fmt.Printf("Pop %d Error: %v\n", i+1, err)
			break
		}
		fmt.Printf("Popped: %d\n", val)
	}

}

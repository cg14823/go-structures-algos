package structures

import "errors"

// stack is a basic data structure for integers.
// It follows a LIFO pattern.
// Not suitable for concurrent use
type stack struct {
	s []int
}

func InitStack() *stack {
	return &stack{make([]int, 0)}
}

func (s *stack) Pop() (int, error){
	if len(s.s) == 0 {
		return 0, errors.New("stack is empty")
	}

	var element int
	element, s.s = s.s[len(s.s) -1], s.s[:len(s.s) - 1]
	return element, nil
}

func (s *stack) Push(element int) {
	s.s = append(s.s, element)
}

func (s *stack) Size() int {
	return len(s.s)
}
package model

import "errors"

//ArrayStack implement stack using slice
type ArrayStack struct {
	arr     []int
	size    int
	maxSize int
}

//NewStack will return stack instance
func NewArrayStack(maxSize int) *ArrayStack {
	arr := make([]int, 0)
	return &ArrayStack{arr: arr, size: 0, maxSize: maxSize}
}

func (s *ArrayStack) Push(item int) error {
	if len(s.arr) == s.maxSize {
		return errors.New("stack overflow")
	}
	s.arr = append(s.arr, item)
	s.size++
	return nil
}

func (s *ArrayStack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("empty stack")
	}
	tmp := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	s.size--
	return tmp, nil
}

func (s *ArrayStack) Top() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("empty stack")
	}
	return s.arr[0], nil
}

func (s *ArrayStack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *ArrayStack) Size() int {
	return s.size
}

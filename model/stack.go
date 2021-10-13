package model

import "errors"

type IntStack struct {
	arr     []int
	Size    int
	maxSize int
}

func NewIntStack(maxSize int) *IntStack {
	arr := make([]int, 0)
	return &IntStack{arr: arr, Size: 0, maxSize: maxSize}
}

func (s *IntStack) Push(item int) error {
	if len(s.arr) == s.maxSize {
		return errors.New("stack overflow")
	}
	s.arr = append(s.arr, item)
	s.Size++
	return nil
}

func (s *IntStack) Pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, errors.New("empty stack")
	}
	tmp := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	s.Size--
	return tmp, nil
}

func (s *IntStack) Top() (int, error) {
	if len(s.arr) == 0 {
		return 0, errors.New("empty stack")
	}
	return s.arr[0], nil
}

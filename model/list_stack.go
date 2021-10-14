package model

import "errors"

type ListStack struct {
	top  *ListNode
	size int
}

func NewListStack() *ListStack {
	return &ListStack{size: 0}
}

func (l *ListStack) isEmpty() bool {
	return l.size == 0
}

func (l *ListStack) Size() int {
	return l.size
}

func (l *ListStack) Push(val int) {
	if l.isEmpty() {
		top := &ListNode{
			Val: val,
		}
		l.top = top
		l.size++
	}
	newTop := &ListNode{
		Val: val,
	}
	newTop.Next = l.top
}

func (l *ListStack) Pop() (int, error) {
	if l.isEmpty() {
		return 0, errors.New("empty stack")
	}
	currentTop := l.top
	l.top = l.top.Next
	return currentTop.Val, nil
}

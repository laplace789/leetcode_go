package model

import (
	"errors"
	"testing"
)

func Test_listStack_Push(t *testing.T) {
	stack1 := NewListStack()

	stack2 := NewListStack()
	stack2.Push(1)

	tests := []struct {
		name    string
		stack   *ListStack
		wantLen int
	}{
		{
			name:    "tes1 push empty stack",
			stack:   stack1,
			wantLen: 1,
		},
		{
			name:    "tes2 push 1 val to stack",
			stack:   stack2,
			wantLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Push(5)
			got := tt.wantLen
			if got != tt.wantLen {
				t.Errorf("want len = %v ,got len = %v", tt.wantLen, got)
			}
		})
	}
}

func Test_ListStack_Pop(t *testing.T) {
	stack1 := NewListStack()

	stack2 := NewListStack()
	stack2.Push(1)
	stack2.Push(2)

	tests := []struct {
		name    string
		stack   *ListStack
		wantLen int
		wantVal int
		wantErr error
	}{
		{
			name:    "tes1 pop empty stack",
			stack:   stack1,
			wantLen: 0,
			wantVal: 0,
			wantErr: errors.New("empty stack"),
		},
		{
			name:    "tes2 pop 1 val to stack",
			stack:   stack2,
			wantLen: 1,
			wantVal: 2,
			wantErr: errors.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotErr := tt.stack.Pop()
			gotLen := tt.stack.Size()

			var gotErrorStr string
			if gotErr == nil {
				gotErrorStr = ""
			} else {
				gotErrorStr = gotErr.Error()
			}

			if gotLen != tt.wantLen {
				t.Errorf("want len = %v ,got len = %v", tt.wantLen, gotLen)
			}

			if gotVal != tt.wantVal {
				t.Errorf("want val = %v ,got val = %v", tt.wantVal, gotVal)
			}

			if gotErrorStr != tt.wantErr.Error() {
				t.Errorf("want err = %v ,got err = %v", tt.wantLen, gotErrorStr)
			}
		})
	}
}

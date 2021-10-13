package model

import (
	"errors"
	"testing"
)

func Test_IntStackPush(t *testing.T) {
	type arg struct {
		stacks []*IntStack
	}

	args := &arg{}
	stackArr := make([]*IntStack, 0)
	stack1 := NewIntStack(5)
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)

	stack2 := NewIntStack(3)
	stack2.Push(1)
	stack2.Push(2)
	stack2.Push(3)

	stackArr = append(stackArr, stack1)
	stackArr = append(stackArr, stack2)

	args.stacks = stackArr

	tests := []struct {
		name    string
		stack   *IntStack
		wantLen int
		wantErr error
	}{
		{
			name:    "tes1 normal push",
			stack:   stack1,
			wantLen: 4,
			wantErr: errors.New(""),
		},
		{
			name:    "tes2 push over stack size",
			stack:   stack2,
			wantLen: 3,
			wantErr: errors.New("stack overflow"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotErrorStr string
			err := tt.stack.Push(1)
			if err == nil {
				gotErrorStr = ""
			} else {
				gotErrorStr = err.Error()
			}

			if gotErrorStr != tt.wantErr.Error() {
				t.Errorf("wanrErr = %v ,getErr = %v", tt.wantErr, gotErrorStr)
			}
			if tt.wantLen != tt.stack.Size {
				t.Errorf("wanrLen = %v ,getLen = %v", tt.wantLen, tt.stack.Size)
			}
		})
	}

}

func Test_IntStackPop(t *testing.T) {
	type arg struct {
		stacks []*IntStack
	}

	args := &arg{}
	stackArr := make([]*IntStack, 0)
	stack1 := NewIntStack(5)
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)

	stack2 := NewIntStack(3)

	stackArr = append(stackArr, stack1)
	stackArr = append(stackArr, stack2)

	args.stacks = stackArr

	tests := []struct {
		name    string
		stack   *IntStack
		wantLen int
		wantVal int
		wantErr error
	}{
		{
			name:    "tes1 pop 1 value",
			stack:   stack1,
			wantLen: 2,
			wantVal: 3,
			wantErr: errors.New(""),
		},
		{
			name:    "tes2 pop empty stack",
			stack:   stack2,
			wantLen: 0,
			wantVal: 0,
			wantErr: errors.New("empty stack"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, err := tt.stack.Pop()
			var gotErrorStr string
			if err == nil {
				gotErrorStr = ""
			} else {
				gotErrorStr = err.Error()
			}
			if gotErrorStr != tt.wantErr.Error() {
				t.Errorf("wanrErr = %v ,gotErr = %v", tt.wantErr, gotErrorStr)
			}
			if tt.wantLen != tt.stack.Size {
				t.Errorf("wanrLen = %v ,gotLen = %v", tt.wantLen, tt.stack.Size)
			}

			if tt.wantVal != gotVal {
				t.Errorf("wanrVal = %v ,gotVal = %v", tt.wantVal, gotVal)
			}
		})
	}
}

func Test_IntStackPop_multi(t *testing.T) {
	type arg struct {
		stacks []*IntStack
	}

	args := &arg{}
	stackArr := make([]*IntStack, 0)
	stack1 := NewIntStack(5)
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)

	stackArr = append(stackArr, stack1)

	args.stacks = stackArr

	tests := []struct {
		name     string
		stack    *IntStack
		want1Len int
		want1Val int
		want1Err error
		want2Len int
		want2Val int
		want2Err error
	}{
		{
			name:     "tes1 pop twice value",
			stack:    stack1,
			want1Len: 2,
			want1Val: 3,
			want1Err: errors.New(""),
			want2Len: 1,
			want2Val: 2,
			want2Err: errors.New(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1Val, errPop1 := tt.stack.Pop()
			var gotErrorStr1 string
			if errPop1 == nil {
				gotErrorStr1 = ""
			} else {
				gotErrorStr1 = errPop1.Error()
			}
			if gotErrorStr1 != tt.want1Err.Error() {
				t.Errorf("want1Err = %v ,got1Err = %v", tt.want1Err, gotErrorStr1)
			}
			if tt.want1Len != tt.stack.Size {
				t.Errorf("wanr1Len = %v ,got1Len = %v", tt.want1Len, tt.stack.Size)
			}

			if tt.want1Val != got1Val {
				t.Errorf("wanr1Val = %v ,got1Val = %v", tt.want1Val, got1Val)
			}

			got2Val, errPop2 := tt.stack.Pop()
			var gotErrorStr2 string
			if errPop2 == nil {
				gotErrorStr2 = ""
			} else {
				gotErrorStr2 = errPop2.Error()
			}
			if gotErrorStr2 != tt.want2Err.Error() {
				t.Errorf("want2Err = %v ,got2Err = %v", tt.want2Err, gotErrorStr2)
			}
			if tt.want2Len != tt.stack.Size {
				t.Errorf("wanr2Len = %v ,got2Len = %v", tt.want2Len, tt.stack.Size)
			}

			if tt.want2Val != got2Val {
				t.Errorf("want2Val = %v ,got2Val = %v", tt.want2Val, got2Val)
			}
		})
	}
}

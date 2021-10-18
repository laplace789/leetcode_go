package model

import (
	"testing"
)

func Test_Int2ListNOdeWithCycle(t *testing.T) {
	type arg struct {
		arr            []int
		cycleHeadIndex int
	}

	tests := []struct {
		name    string
		args    arg
		wantVal int
	}{
		{
			name:    "tes1 form cycle",
			args:    arg{arr: []int{3, 2, 0, -4}, cycleHeadIndex: 1},
			wantVal: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := Int2ListNOdeWithCycle(tt.args.arr, tt.args.cycleHeadIndex)
			for i := 0; i < len(tt.args.arr); i++ {
				//last node
				if i == len(tt.args.arr)-1 {
					if node.Next == nil {
						t.Errorf("for, cycle fail,cycle should't contain Next with nil")
					}
					got := node.Next.Val
					if got != tt.wantVal {
						t.Errorf("want = %v ,got val = %v", tt.wantVal, got)
					}
					return
				}
				node = node.Next
			}
		})
	}
}

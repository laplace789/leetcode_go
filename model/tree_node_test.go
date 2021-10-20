package model

import (
	"testing"
)

func Test_Int2TreeNode(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	tests := []struct {
		name         string
		arr          []int
		wantLeftVal  int
		wantRightVal int
	}{
		{
			name: "tes1 pop empty stack",
			arr:  arr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := Int2TreeNode(tt.arr)
			gotLeftVal := tree.Left.Val

			if tt.wantLeftVal != gotLeftVal {
				t.Errorf("wantLeftVal = %v,gotLeftVal = %v", tt.wantLeftVal, gotLeftVal)
			}
		})
	}
}

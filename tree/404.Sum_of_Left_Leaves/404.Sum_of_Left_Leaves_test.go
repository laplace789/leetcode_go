package tree

import (
	"leetcode_go/model"
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	root1 := model.Int2TreeNode([]int{3, 9, 20, model.NULL, model.NULL, 15, 7})
	root2 := model.Int2TreeNode([]int{1, 2, 3, 4, 5})
	tests := []struct {
		name    string
		root    *model.TreeNode
		wantVal int
	}{
		{
			name:    "tes1",
			root:    root1,
			wantVal: 24,
		},
		{
			name:    "tes2",
			root:    root2,
			wantVal: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumOfLeftLeaves(tt.root)
			if got != tt.wantVal {
				t.Errorf("wantVal = %v , got = %v",tt.wantVal,got)
			}
		})
	}
}

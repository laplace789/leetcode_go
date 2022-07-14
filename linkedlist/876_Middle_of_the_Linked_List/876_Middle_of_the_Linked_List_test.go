package linkedlist

import (
	"leetcode_go/model"
	"testing"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  []int
	}{
		{
			name:  "test1",
			nodes: []int{1, 2, 3, 4, 5, 6},
			want:  []int{4, 5, 6},
		},
		{
			name:  "test2",
			nodes: []int{1, 2, 3, 4, 5},
			want:  []int{3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := model.Int2ListNOde(tt.nodes)
			want := model.Int2ListNOde(tt.want)
			got := middleNode(node)
			for want != nil && got != nil {
				if got.Val != want.Val {
					t.Errorf("got val = %v,want val = %v", got.Next, want.Next)
				}
				want = want.Next
				got = got.Next
			}
		})
	}
}

package linkedlist

import (
	"leetcode_go/model"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	node1 := model.Int2ListNOde([]int{1, 1, 1, 1})
	node2 := model.Int2ListNOdeWithCycle([]int{3, 2, 0, -4}, 1)

	tests := []struct {
		name    string
		list    *ListNode
		wantVal bool
	}{
		{
			name:    "test1 no cycle",
			list:    node1,
			wantVal: false,
		},
		{
			name:    "tes2 with cycle",
			list:    node2,
			wantVal: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasCycle(tt.list)
			if got != tt.wantVal {
				t.Errorf("want val = %v,got val = %v", tt.wantVal, got)
			}
		})
	}
}

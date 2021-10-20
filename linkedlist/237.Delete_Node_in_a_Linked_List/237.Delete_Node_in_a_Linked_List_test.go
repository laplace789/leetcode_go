package linkedlist

import (
	"leetcode_go/model"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	node1 := model.Int2ListNOde([]int{4, 5, 1, 9})

	tests := []struct {
		name      string
		list      *ListNode
		deletePos int
		wantVal   []int
		wantLen   int
	}{
		{
			name:      "test1 no cycle",
			list:      node1,
			deletePos: 1,
			wantVal:   []int{4, 1, 9},
			wantLen:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			next := tt.list
			for i := 0; i < tt.deletePos; i++ {
				next = tt.list.Next
			}
			deleteNode(next)
			//trans list to arr
			gotArr := model.ListNode2Int(tt.list)
			if len(gotArr) != tt.wantLen {
				t.Errorf("list len is not equal after delete wantLen = %v ,gotLen = %v", tt.wantLen, len(gotArr))
			}

			for i := 0; i < len(gotArr); i++ {
				if gotArr[i] != tt.wantVal[i] {
					t.Errorf(" wantVal = %v ,gotVal = %v", tt.wantVal, gotArr[i])
				}
			}
		})
	}
}

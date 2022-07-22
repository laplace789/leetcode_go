package linkedlist

import (
	"leetcode_go/model"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {

	tests := []struct {
		name    string
		arr     []int
		n       int
		wantArr []int
	}{
		{
			name:    "test1",
			arr:     []int{1, 2, 3, 4, 5},
			n:       2,
			wantArr: []int{1, 2, 3, 5},
		},
		{
			name:    "test2",
			arr:     []int{1, 2},
			n:       1,
			wantArr: []int{1},
		},
		{
			name:    "test3",
			arr:     []int{1},
			n:       1,
			wantArr: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := model.Int2ListNOde(tt.arr)
			gotNode := removeNthFromEndTwo(head, tt.n)
			got := model.ListNode2Int(gotNode)

			if len(got) == len(tt.wantArr) {
				for i := 0; i < len(got); i++ {
					if got[i] != tt.wantArr[i] {
						t.Errorf("want  = %v,got  = %v", tt.wantArr[i], got[i])
					}
				}
			} else {
				t.Errorf("want len = %v,got len = %v", len(tt.wantArr), len(got))
			}

		})
	}
}

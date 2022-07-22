package linkedlist

import (
	"leetcode_go/model"
	"testing"
)

func Test_swapPairs(t *testing.T) {

	tests := []struct {
		name    string
		arr     []int
		wantArr []int
	}{
		{
			name:    "test1",
			arr:     []int{1, 2, 3, 4},
			wantArr: []int{2, 1, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := model.Int2ListNOde(tt.arr)
			gotNode := swapPairs(head)

			t.Logf("test = %v", gotNode)
		})
	}
}

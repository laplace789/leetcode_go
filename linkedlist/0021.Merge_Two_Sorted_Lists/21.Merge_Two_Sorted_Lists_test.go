package linkedList

import (
	"leetcode_go/model"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tes1",
			args: args{arr1: []int{1, 2, 4}, arr2: []int{1, 3, 4}},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "tes2",
			args: args{arr1: []int{9, 9, 9, 9, 9}, arr2: []int{1}},
			want: []int{1, 9, 9, 9, 9, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := model.Int2ListNOde(tt.args.arr1)
			l2 := model.Int2ListNOde(tt.args.arr2)
			got := mergeTwoLists(l1, l2)
			gotArr := model.ListNode2Int(got)
			if len(gotArr) != len(tt.want) {
				t.Errorf("len size not equal")
			}

			for i := 0; i < len(gotArr); i++ {
				if gotArr[i] != tt.want[i] {
					t.Errorf("want value = %d got value = %d", tt.want[i], gotArr[i])
				}
			}
		})
	}
}

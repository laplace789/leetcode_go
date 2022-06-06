package tree

import (
	"leetcode_go/model"
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				root: model.Int2TreeNode([]int{1, model.NULL, 2, 3}),
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test2",
			args: args{
				root: model.Int2TreeNode([]int{}),
			},
			want: []int{},
		},
		{
			name: "test2",
			args: args{
				root: model.Int2TreeNode([]int{1}),
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

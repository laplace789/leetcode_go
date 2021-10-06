package array

import (
	"testing"
)

func Test_numIdenticalPairs(t *testing.T){
	type args struct {
		arr1 []int
	}

	tests := []struct {
		name    string
		args    args
		want   int
	}{
		{
			name: "tes1",
			args: args{arr1: []int{1,2,3,1,1,3}},
			want: 4,
		},
		{
			name: "tes2",
			args: args{arr1: []int{1,1,1,1}},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := numIdenticalPairs(tt.args.arr1)
			if res != tt.want{
				t.Errorf("want value = %d got value = %d",tt.want, res)
			}
		})
	}
}

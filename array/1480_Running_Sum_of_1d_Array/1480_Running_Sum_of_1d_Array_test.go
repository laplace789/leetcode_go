package array

import "testing"

func Test_runningSum(t *testing.T) {
	type args struct {
		arr1 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tes1",
			args: args{arr1: []int{1, 1, 1, 1, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "tes2",
			args: args{arr1: []int{1, 2, 3, 4}},
			want: []int{1, 3, 6, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := runningSum(tt.args.arr1)
			if len(res) != len(tt.want) {
				t.Errorf("len different, want len = %d, get len = %d", len(tt.want), len(res))
			}
			for i := 0; i < len(res); i++ {
				if res[i] != tt.want[i] {
					t.Errorf("want value = %d got value = %d", tt.want[i], res[i])
				}
			}
		})
	}
}

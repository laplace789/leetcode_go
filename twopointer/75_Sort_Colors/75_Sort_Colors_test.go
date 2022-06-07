package twopointer

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantArr []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			wantArr: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, 0, 1},
			},
			wantArr: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if len(tt.wantArr) != len(tt.args.nums) {
				t.Errorf("want len = %v,got len = %v", tt.wantArr, tt.args.nums)
			}

			for i := 0; i < len(tt.wantArr); i++ {
				if tt.wantArr[i] != tt.args.nums[i] {
					t.Errorf("want num = %v,got num = %v", tt.wantArr[i], tt.args.nums[i])
				}
			}
		})
	}
}

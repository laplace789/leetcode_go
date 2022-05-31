package sort

import "testing"

func Test_largestNumber(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantStr string
	}{
		{
			name:    "test1",
			nums:    []int{10, 2},
			wantStr: "210",
		},
		{
			name:    "test2",
			nums:    []int{3, 30},
			wantStr: "330",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestNumber(tt.nums)

			if got != tt.wantStr {
				t.Errorf("want str = %s,got str = %v", got, tt.wantStr)
			}
		})
	}
}

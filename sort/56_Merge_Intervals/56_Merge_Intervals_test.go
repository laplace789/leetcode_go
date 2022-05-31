package sort

import "testing"

func Test_merge(t *testing.T) {
	tests := []struct {
		name    string
		data    [][]int
		wantRes [][]int
	}{
		{
			name: "test1",
			data: [][]int{
				{1, 3}, {2, 6}, {8, 10}, {15, 18},
			},
			wantRes: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.data)
			if len(got) != len(tt.wantRes) {
				t.Errorf("want len = %d,got len = %d", len(tt.wantRes), len(got))
			}

			for i := 0; i < len(got); i++ {
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] != tt.wantRes[i][j] {
						t.Errorf("want data = %d,got data = %d, i =%d j=%d", tt.wantRes[i][j], got[i][j], i, j)
					}
				}
			}
		})
	}
}

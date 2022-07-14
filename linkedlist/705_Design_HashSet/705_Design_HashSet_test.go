package linkedlist

import (
	"testing"
)

func TestMyHashSet_AllOp(t *testing.T) {
	//A -> Add,D->Delete,C->Contains

	tests := []struct {
		name     string
		keys     []int
		ops      []string
		checkArr []int
		wantBool []bool
	}{
		{
			name:     "test1",
			keys:     []int{1, 1, 1, 4, 5, 1, 8, 9, 10},
			ops:      []string{"A", "A", "D", "A", "A", "A", "D", "A", "D"}, //define add and remove op
			checkArr: []int{1, 4, 5, 9, 888},                                //use contains to check correction
			wantBool: []bool{true, true, true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := Constructor()
			for i := 0; i < len(tt.keys); i++ {
				switch tt.ops[i] {
				case "A":
					{
						set.Add(tt.keys[i])
					}
				case "D":
					{
						set.Remove(tt.keys[i])
					}
				}
			}
			for i := 0; i < len(tt.checkArr); i++ {
				got := set.Contains(tt.checkArr[i])
				if got != tt.wantBool[i] {
					t.Errorf("val[i] = %v got res = %v,want val = %v", tt.checkArr[i], got, tt.wantBool[i])
				}
			}
		})
	}
}

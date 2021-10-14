package string

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{
			name:   "tes1",
			jewels: "aA",
			stones: "aAAbbbb",
			want:   3,
		},
		{
			name:   "tes2",
			jewels: "z",
			stones: "ZZ",
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numJewelsInStones(tt.jewels, tt.stones)
			if got != tt.want {
				t.Errorf("want val = %v ,got val = %v",tt.want,got)
			}
		})
	}
}

package string

import "testing"

func Test_countSegments(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		want int
	}{
		{
			name: "tes1",
			arg1: "Hello, my name is John",
			want: 5,
		},
		{
			name: "tes2",
			arg1: "Hello",
			want: 1,
		},
		{
			name: "tes3",
			arg1: "love live! mu'sic forever",
			want: 4,
		},
		{
			name: "tes4",
			arg1: "",
			want: 0,
		},
		{
			name: "tes5",
			arg1: "              ",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSegments(tt.arg1)
			if got != tt.want {
				t.Errorf("want value = %v got value = %v", tt.want, got)
			}
		})
	}

}

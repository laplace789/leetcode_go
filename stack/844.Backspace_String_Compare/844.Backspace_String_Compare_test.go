package stack

import "testing"

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want bool
	}{
		{
			name: "tes1",
			arg1: "ab#c",
			arg2: "ad#c",
			want: true,
		},
		{
			name: "tes2",
			arg1: "ab##",
			arg2: "c#d#",
			want: true,
		},
		{
			name: "tes3",
			arg1: "xywrrmp",
			arg2: "xywrrmu#p",
			want: true,
		},
		{
			name: "tes3",
			arg1: "#",
			arg2: "a#p",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := backspaceCompare(tt.arg1, tt.arg2)
			if got != tt.want {
				t.Errorf("want value = %v got value = %v", tt.want, got)
			}
		})
	}
}

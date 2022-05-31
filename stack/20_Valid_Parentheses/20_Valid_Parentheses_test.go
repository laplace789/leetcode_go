package stack

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "tes1",
			args: "()",
			want: true,
		},
		{
			name: "tes2",
			args: "()[]{}",
			want: true,
		},
		{
			name: "test3",
			args: "{]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.args)
			if got != tt.want {
				t.Errorf("want value = %v got value = %v", tt.want, got)
			}
		})
	}
}

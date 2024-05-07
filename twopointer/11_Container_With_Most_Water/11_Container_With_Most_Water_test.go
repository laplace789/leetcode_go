package twopointer

import (
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test1",
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainerWithMostWater(tt.input)
			if got != tt.want {
				t.Errorf("got = %d want = %d", got, tt.want)
			}
		})
	}
}

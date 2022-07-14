package linkedlist

import (
	"leetcode_go/model"
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  int
	}{
		{
			name:  "test1",
			nodes: []int{1, 0, 1},
			want:  5,
		},
		{
			name:  "test2",
			nodes: []int{1, 0, 1, 0},
			want:  10,
		},
		{
			name:  "test3",
			nodes: []int{1},
			want:  1,
		},
		{
			name:  "test4",
			nodes: []int{0},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := model.Int2ListNOde(tt.nodes)
			got := getDecimalValue(node)
			if got != tt.want {
				t.Errorf("got val = %v,want val = %v", got, tt.want)
			}
		})
	}
}

func Test_getDecimalValueHigh(t *testing.T) {
	tests := []struct {
		name  string
		nodes []int
		want  int
	}{
		{
			name:  "test1",
			nodes: []int{1, 0, 1},
			want:  5,
		},
		{
			name:  "test2",
			nodes: []int{1, 0, 1, 0},
			want:  10,
		},
		{
			name:  "test3",
			nodes: []int{1},
			want:  1,
		},
		{
			name:  "test4",
			nodes: []int{0},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := model.Int2ListNOde(tt.nodes)
			got := getDecimalValueHigh(node)
			if got != tt.want {
				t.Errorf("got val = %v,want val = %v", got, tt.want)
			}
		})
	}
}

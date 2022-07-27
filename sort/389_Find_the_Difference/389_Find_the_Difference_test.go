package sort

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "test1",
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: uint8(101),
		},
		{
			name: "test2",
			args: args{
				s: "",
				t: "y",
			},
			want: uint8(121),
		},
		{
			name: "test3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: uint8(97),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

package internal

import "testing"

func Test_diff(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"fail", args{2, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diff(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

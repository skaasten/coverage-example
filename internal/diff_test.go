package internal

import (
	"math/rand"
	"testing"
	"time"
)

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
		{"pass", args{3, 2}, rand.Intn(2)},
	}
	for _, tt := range tests {
		rand.Seed(time.Now().UnixNano())
		t.Run(tt.name, func(t *testing.T) {
			if got := diff(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

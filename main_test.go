package main

import (
	"testing"
)

func Test_sum(t *testing.T) {
	expected := 4
	got := sum(1, 3)
	if expected != got {
		t.Errorf("sum() = %v, want %v", got, expected)
	}
}

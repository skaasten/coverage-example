package internal

import (
	"math/rand"
	"testing"
	"time"
)

func Test_diff(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	expected := rand.Intn(2)
	got := diff(3, 2)
	if got != expected {
		t.Errorf("diff() = %v, want %v", got, expected)
	}
}

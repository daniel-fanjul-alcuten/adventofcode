package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(7, 3, S); n != 6 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(50, 6, X); n != 121 {
		t.Error(n)
	}
}

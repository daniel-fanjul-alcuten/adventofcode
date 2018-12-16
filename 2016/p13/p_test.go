package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(10, Position{7, 4}); n != 11 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(1358, Position{31, 39}); n != 96 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P2(1358, 50); n != 141 {
		t.Error(n)
	}
}

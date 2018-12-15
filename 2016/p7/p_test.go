package p

import (
	"testing"
)

func TestP1_S1(t *testing.T) {
	if n := P1(S1); n != 2 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X); n != 110 {
		t.Error(n)
	}
}

func TestP2_S2(t *testing.T) {
	if n := P2(S2); n != 3 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P2(X); n != 242 {
		t.Error(n)
	}
}

package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S); n != 42 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X); n != 318117 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	t.Parallel()
	if n := P2(X); n != 9227771 {
		t.Error(n)
	}
}

package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S); n != 1 {
		t.Error(n)
	}
}

func TestP1_X1(t *testing.T) {
	if n := P1(X1); n != 636 {
		t.Error(n)
	}
}

func TestP2_X2(t *testing.T) {
	if n := P2(X1, X2); n != 674 {
		t.Error(n)
	}
}

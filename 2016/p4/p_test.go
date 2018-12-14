package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S); n != 1514 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X); n != 158835 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P2(S, "tttttt"); n != 123 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P2(X, "northpole"); n != 993 {
		t.Error(n)
	}
}

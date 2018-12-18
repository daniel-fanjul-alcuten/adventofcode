package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S); n != 1 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X); n != 1050 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P1(P2(S)); n != 2 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P1(P2(X)); n != 1921 {
		t.Error(n)
	}
}

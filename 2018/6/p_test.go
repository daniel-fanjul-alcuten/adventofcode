package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S); n != 17 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X); n != 5035 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P2(32, S); n != 16 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P2(10000, X); n != 35294 {
		t.Error(n)
	}
}

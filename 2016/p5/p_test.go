package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	t.Parallel()
	if p := P1(S); p != "18f47a30" {
		t.Error(p)
	}
}

func TestP1_X(t *testing.T) {
	t.Parallel()
	if p := P1(X); p != "d4cd2ee1" {
		t.Error(p)
	}
}

func TestP2_S(t *testing.T) {
	t.Parallel()
	if p := P2(S); p != "05ace8e3" {
		t.Error(p)
	}
}

func TestP2_X(t *testing.T) {
	t.Parallel()
	if p := P2(X); p != "f2c730e5" {
		t.Error(p)
	}
}

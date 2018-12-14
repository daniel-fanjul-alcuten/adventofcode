package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if s := P1(S); s != "1985" {
		t.Error(s)
	}
}

func TestP1_X(t *testing.T) {
	if s := P1(X); s != "56855" {
		t.Error(s)
	}
}

func TestP2_S(t *testing.T) {
	if s := P2(S); s != "5DB3" {
		t.Error(s)
	}
}

func TestP2_X(t *testing.T) {
	if s := P2(X); s != "B3C27" {
		t.Error(s)
	}
}

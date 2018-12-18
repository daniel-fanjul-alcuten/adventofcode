package p

import (
	"testing"
)

func TestP1(t *testing.T) {
	if n := P1(5); n != 3 {
		t.Error(n)
	}
	if n := P1(3012210); n != 1830117 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P2(5); n != 2 {
		t.Error(n)
	}
	if n := P2(3012210); n != 1417887 {
		t.Error(n)
	}
}

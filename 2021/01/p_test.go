package p

import (
	"testing"
)

func TestP1_S1(t *testing.T) {
	if r := P1(S1...); r != 7 {
		t.Error(r)
	}
}

func TestP1_X1(t *testing.T) {
	if r := P1(X1...); r != 1466 {
		t.Error(r)
	}
}

func TestP2_S1(t *testing.T) {
	if r := P2(S1...); r != 5 {
		t.Error(r)
	}
}

func TestP2_X1(t *testing.T) {
	if r := P2(X1...); r != 1491 {
		t.Error(r)
	}
}

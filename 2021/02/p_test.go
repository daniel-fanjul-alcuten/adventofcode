package p

import (
	"testing"
)

func TestP1_S1(t *testing.T) {
	if f, d := P1(S1...); f*d != 150 {
		t.Error(f * d)
	}
}

func TestP1_X1(t *testing.T) {
	if f, d := P1(X1...); f*d != 1989014 {
		t.Error(f * d)
	}
}

func TestP2_S1(t *testing.T) {
	if f, d := P2(S1...); f*d != 900 {
		t.Error(f * d)
	}
}

func TestP2_X1(t *testing.T) {
	if f, d := P2(X1...); f*d != 2006917119 {
		t.Error(f * d)
	}
}

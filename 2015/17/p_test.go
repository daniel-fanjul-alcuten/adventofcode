package p

import "testing"

func TestP1_S(t *testing.T) {
	if n := P1(25, 20, 15, 10, 5, 5); n != 4 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(150, X...); n != 4372 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P2(25, 20, 15, 10, 5, 5); n != 3 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P2(150, X...); n != 4 {
		t.Error(n)
	}
}

var X = []int{11, 30, 47, 31, 32, 36, 3, 1, 5, 3, 32, 36, 15, 11, 46, 26, 28, 1, 19, 3}

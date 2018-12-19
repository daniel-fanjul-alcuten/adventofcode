package p

import "testing"

func TestP1_S(t *testing.T) {
	if n := P1(false, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11); n != 99 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(false, 1, 3, 5, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113); n != 10439961859 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P1(true, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11); n != 44 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P1(true, 1, 3, 5, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113); n != 72050269 {
		t.Error(n)
	}
}

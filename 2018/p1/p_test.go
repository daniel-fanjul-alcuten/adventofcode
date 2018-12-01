package p

import "testing"

func TestP1_Samples(t *testing.T) {
	if r := P1(0, []int{+1, -2, +3, +1}); r != 3 {
		t.Error(r)
	}
	if r := P1(0, []int{+1, +1, +1}); r != 3 {
		t.Error(r)
	}
	if r := P1(0, []int{+1, +1, -2}); r != 0 {
		t.Error(r)
	}
	if r := P1(0, []int{-1, -2, -3}); r != -6 {
		t.Error(r)
	}
}

func TestP1_X(t *testing.T) {
	if r := P1(0, X); r != 400 {
		t.Error(r)
	}
}

func TestP2_Samples(t *testing.T) {
	if r := P2(0, []int{+1, -2, +3, +1}); r != 2 {
		t.Error(r)
	}
	if r := P2(0, []int{+1, -1}); r != 0 {
		t.Error(r)
	}
	if r := P2(0, []int{+3, +3, +4, -2, -4}); r != 10 {
		t.Error(r)
	}
	if r := P2(0, []int{+7, +7, -2, -7, -4}); r != 14 {
		t.Error(r)
	}
}

func TestP2_X(t *testing.T) {
	if r := P2(0, X); r != 232 {
		t.Error(r)
	}
}

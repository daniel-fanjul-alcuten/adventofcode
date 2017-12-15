package p2

import "testing"

func TestP1_5195(t *testing.T) {
	if v := (P1([][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	})); v != 18 {
		t.Error(v)
	}
}

func TestP1_X(t *testing.T) {
	if v := (P1(ParseN(X))); v != 47136 {
		t.Error(v)
	}
}

func TestP2_5195(t *testing.T) {
	if v := (P2([][]int{
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
	})); v != 9 {
		t.Error(v)
	}
}

func TestP2_X(t *testing.T) {
	if v := (P2(ParseN(X))); v != 250 {
		t.Error(v)
	}
}

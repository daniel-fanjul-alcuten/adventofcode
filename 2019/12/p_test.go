package p

import "testing"

func TestP1(t *testing.T) {
	if e := P1(10, pos{-1, 0, 2}, pos{2, -10, -7}, pos{4, -8, 8}, pos{3, 5, -1}); e != 179 {
		t.Error(e)
	}
	if e := P1(100, pos{-8, -10, 0}, pos{5, 5, 10}, pos{2, -7, 3}, pos{9, -8, -3}); e != 1940 {
		t.Error(e)
	}
	if e := P1(1000, X...); e != 8538 {
		t.Error(e)
	}
}

func TestP2(t *testing.T) {
	if n := P2(pos{-1, 0, 2}, pos{2, -10, -7}, pos{4, -8, 8}, pos{3, 5, -1}); n != 2772 {
		t.Error(n)
	}
	if n := P2(pos{-8, -10, 0}, pos{5, 5, 10}, pos{2, -7, 3}, pos{9, -8, -3}); n != 4686774924 {
		t.Error(n)
	}
	if n := P2(X...); n != 506359021038056 {
		t.Error(n)
	}
}

var X = []pos{
	{-9, 10, -1},
	{-14, -8, 14},
	{1, 5, 6},
	{-19, 7, 8},
}

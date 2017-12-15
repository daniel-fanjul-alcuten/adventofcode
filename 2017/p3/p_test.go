package p3

import "testing"

func TestP1(t *testing.T) {
	if v := P1(1); v != 0 {
		t.Error(v)
	}
	if v := P1(12); v != 3 {
		t.Error(v)
	}
	if v := P1(23); v != 2 {
		t.Error(v)
	}
	if v := P1(1024); v != 31 {
		t.Error(v)
	}
	if v := P1(347991); v != 480 {
		t.Error(v)
	}
}

func TestNext2(t *testing.T) {
	prev := map[Position]int{}
	r := FirstRecord2(prev)
	if n := r.Number; n != 1 {
		t.Error(n)
	}
	r = r.Next2(prev)
	if n := r.Number; n != 1 {
		t.Error(n)
	}
	r = r.Next2(prev)
	if n := r.Number; n != 2 {
		t.Error(n)
	}
	r = r.Next2(prev)
	if n := r.Number; n != 4 {
		t.Error(n)
	}
	r = r.Next2(prev)
	if n := r.Number; n != 5 {
		t.Error(n)
	}
}

func TestP2(t *testing.T) {
	if n := P2(347991); n != 349975 {
		t.Error(n)
	}
}

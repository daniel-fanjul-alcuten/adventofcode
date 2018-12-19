package p

import "testing"

func TestP1(t *testing.T) {
	p := Character{50, 0, 0}
	e := Character{71, 10, 0}
	if m := P1(p, e, 500, false); m != 1824 {
		t.Error(m)
	}
}

func TestP2(t *testing.T) {
	p := Character{50, 0, 0}
	e := Character{71, 10, 0}
	if m := P1(p, e, 500, true); m != 1937 {
		t.Error(m)
	}
}

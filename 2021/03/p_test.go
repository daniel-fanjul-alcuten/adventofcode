package p

import (
	"testing"
)

func TestP1_S1(t *testing.T) {
	if g, e := P1(5, S1...); g != 0b_10110 {
		t.Errorf("%05b", g)
	} else if e != 0b_01001 {
		t.Errorf("%05b", e)
	} else if g*e != 198 {
		t.Error(g * e)
	}
}

func TestP1_X1(t *testing.T) {
	if g, e := P1(12, X1...); g != 0b_1010001110 {
		t.Errorf("%05b", g)
	} else if e != 0b_110101110001 {
		t.Errorf("%05b", e)
	} else if g*e != 2250414 {
		t.Error(g * e)
	}
}

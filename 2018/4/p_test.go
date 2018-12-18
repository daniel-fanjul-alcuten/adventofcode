package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if ee, err := Parse(S); err != nil {
		t.Error(err)
	} else if g, m := P1(ee); g != 10 {
		t.Error(g)
	} else if m != 24 {
		t.Error(m)
	}
}

func TestP1_X(t *testing.T) {
	if ee, err := Parse(X); err != nil {
		t.Error(err)
	} else if g, m := P1(ee); g != 1601 {
		t.Error(g)
	} else if m != 46 {
		t.Error(m)
	} else if s := g * m; s != 73646 {
		t.Error(s)
	}
}

func TestP2_S(t *testing.T) {
	if ee, err := Parse(S); err != nil {
		t.Error(err)
	} else if g, m := P2(ee); g != 99 {
		t.Error(g)
	} else if m != 45 {
		t.Error(m)
	}
}

func TestP2_X(t *testing.T) {
	if ee, err := Parse(X); err != nil {
		t.Error(err)
	} else if g, m := P2(ee); g != 163 {
		t.Error(g)
	} else if m != 29 {
		t.Error(m)
	} else if s := g * m; s != 4727 {
		t.Error(s)
	}
}

package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	n, p := Parse(S)
	if l := len(p); l != 0 {
		t.Error(l)
	}
	if s := n.P1(); s != 138 {
		t.Error(s)
	}
}

func TestP1_X(t *testing.T) {
	n, p := Parse(X)
	if l := len(p); l != 0 {
		t.Error(l)
	}
	if s := n.P1(); s != 47647 {
		t.Error(s)
	}
}

func TestP2_S(t *testing.T) {
	n, p := Parse(S)
	if l := len(p); l != 0 {
		t.Error(l)
	}
	if s := n.P2(); s != 66 {
		t.Error(s)
	}
}

func TestP2_X(t *testing.T) {
	n, p := Parse(X)
	if l := len(p); l != 0 {
		t.Error(l)
	}
	if s := n.P2(); s != 23636 {
		t.Error(s)
	}
}
